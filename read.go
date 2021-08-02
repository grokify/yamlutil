package yamlplus

import (
	"log"
	"os"
	"regexp"

	ghoyaml "github.com/ghodss/yaml"
	"gopkg.in/yaml.v3"
)

var rxYamlExtension = regexp.MustCompile(`(?i)\.ya?ml\s*$`)

func ReadFile(filename string) (*yaml.Node, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var node yaml.Node
	err = yaml.Unmarshal(bytes, &node)
	return &node, err
}

// ReadFileAsJson reads a JSON or YAML file as JSON,
// converting YAML to JSON so it can be parsed using
// JSON parsers.
func ReadFileAsJson(filename string) ([]byte, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if rxYamlExtension.MatchString(filename) {
		bytes, err = ghoyaml.YAMLToJSON(bytes)
	}
	return bytes, err
}
