package yamlplus

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	TagBoolean  = "!!bool"
	TagInteger  = "!!int"
	TagMap      = "!!map"
	TagSequence = "!!seq"
	TagString   = "!!str"
)

func unshift(s []string) (string, []string) {
	if len(s) > 0 {
		return s[0], s[1:]
	}
	return "", []string{}
}

func logNode(node *yaml.Node, prefix string) {
	fmt.Printf("%s at LINE [%d] COL [%d] TAG [%s] VAL [%s]\n", prefix, node.Line, node.Column, node.Tag, node.Value)
}

func GetNodeJsonSchemaPath(node *yaml.Node, jsonSchemaPath ...string) (*yaml.Node, error) {
	if node == nil {
		return nil, errors.New("node cannot be nil")
	} else if len(jsonSchemaPath) == 0 {
		return nil, errors.New("no search path")
	}
	curPathPart := jsonSchemaPath[0]

	// logNode(node, "PROC_NODE")

	if node.Value == curPathPart &&
		len(jsonSchemaPath) == 0 {
		return node, nil
	}

	if node.Tag == TagSequence {
		nodeValueInt, err := strconv.Atoi(curPathPart)
		if err != nil {
			panic(err)
		}
		if nodeValueInt < 0 || nodeValueInt >= len(node.Content) {
			return nil, fmt.Errorf("sequence out of range idx [%d] len [%d]", nodeValueInt, len(node.Content))
		}
		if len(jsonSchemaPath) == 1 {
			return node.Content[nodeValueInt], nil
		}
		return GetNodeJsonSchemaPath(
			node.Content[nodeValueInt],
			jsonSchemaPath[1:]...)
	}

	for idxChildNode, childNode := range node.Content {
		// logNode(childNode, "PROC_NODE_CHILD")
		if childNode.Tag == TagString {
			if childNode.Value == curPathPart {
				if len(jsonSchemaPath) == 1 {
					return childNode, nil
				}
				if idxChildNode < len(node.Content)-1 {
					return GetNodeJsonSchemaPath(
						node.Content[idxChildNode+1],
						jsonSchemaPath[1:]...)
				}
				return nil, fmt.Errorf("no next part for [%s]",
					strings.Join(jsonSchemaPath[1:], "/"))
			}
		} else if childNode.Tag == TagMap {
			for idxGrandchildNode, grandchildNode := range childNode.Content {
				// logNode(grandchildNode, "PROC_NODE_GRANDCHILD")
				if grandchildNode.Tag == TagString &&
					grandchildNode.Value == curPathPart {
					// no more paths to check
					if len(jsonSchemaPath) == 1 {
						return grandchildNode, nil
					} else { // more path to check.
						// fmtutil.PrintJSON(grandchildNode)
						// fmt.Printf("SEARCH PATH [%s]\n", strings.Join(jsonSchemaPath[1:], "/"))
						if idxGrandchildNode < len(childNode.Content)-1 {
							// fmt.Printf("FINDING_SEARCH_PATH [%s]\n",
							//	strings.Join(jsonSchemaPath[1:], "/"))
							return GetNodeJsonSchemaPath(
								childNode.Content[idxGrandchildNode+1],
								jsonSchemaPath[1:]...)
						}
						return nil, fmt.Errorf("no next part for [%s]",
							strings.Join(jsonSchemaPath[1:], "/"))
					}
				}
			}
		}

	}
	return nil, fmt.Errorf("nodeName not found [%s]", strings.Join(jsonSchemaPath, ","))
}

/*
// GetNodeJsonSchemaPath returns a node given a JSON Schema Pointer path.
func GetNodeJsonSchemaPath(node *yaml.Node, jsonSchemaPath ...string) (*yaml.Node, error) {
	if node == nil {
		return nil, errors.New("node cannot be nil")
	} else if len(jsonSchemaPath) == 0 {
		return nil, errors.New("no search path")
	}
	curPathPart, jsonSchemaPath := unshift(jsonSchemaPath)

	for i, childNode := range node.Content {
		if childNode.Tag == TagString {
			if childNode.Value == curPathPart {
				if len(jsonSchemaPath) == 0 {
					return childNode, nil
				}
				if i == len(node.Content)-1 {
					return childNode, fmt.Errorf("out of nodes for path [%s]", jsonSchemaPath[0])
				}
				nextChildNode := node.Content[i+1]
				nextPathPart := jsonSchemaPath[0]
				if nextChildNode.Tag == TagMap {
					for _, nextChildNodeMapGrandchildNode := range nextChildNode.Content {
						if nextChildNodeMapGrandchildNode.Value == nextPathPart {
							if len(jsonSchemaPath) == 1 {
								return nextChildNodeMapGrandchildNode, nil
							}
							return GetNodeJsonSchemaPath(nextChildNodeMapGrandchildNode, jsonSchemaPath[1:]...)
						}
					}
				} else if nextChildNode.Tag == TagSequence {
					nextPathPartInt, err := strconv.Atoi(nextPathPart)
					if err != nil {
						return childNode, err
					} else if nextPathPartInt < 0 || nextPathPartInt >= len(nextChildNode.Content) {
						return childNode, fmt.Errorf("index out of range [%d]\n", nextPathPartInt)
					}
					if len(jsonSchemaPath) == 1 {
						return childNode.Content[nextPathPartInt], nil
					}
					return GetNodeJsonSchemaPath(childNode.Content[nextPathPartInt], jsonSchemaPath[1:]...)
				} else {
					return childNode, fmt.Errorf("child not found for [%s/%s]", curPathPart, jsonSchemaPath[0])
				}
			}
		}
	}
	return nil, fmt.Errorf("nodeName not found [%s]", jsonSchemaPath[0])
}
*/
