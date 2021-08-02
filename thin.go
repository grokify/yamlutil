package yamlplus

import "gopkg.in/yaml.v3"

type NodeThin struct {
	Tag     string
	Value   string
	Content []*NodeThin
}

// NodeToThin is used for debugging.
func NodeToThin(node *yaml.Node) *NodeThin {
	if node == nil {
		return nil
	}
	nodeThin := &NodeThin{
		Tag:     node.Tag,
		Value:   node.Value,
		Content: []*NodeThin{}}
	for _, childNode := range node.Content {
		if childNode == nil {
			continue
		}
		childNodeThin := NodeToThin(childNode)
		if childNodeThin != nil {
			nodeThin.Content = append(nodeThin.Content, childNodeThin)
		}
	}
	return nodeThin
}
