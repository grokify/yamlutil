package yamlplus

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFile(filename string) (*yaml.Node, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var node yaml.Node
	err = yaml.Unmarshal(bytes, &node)
	return &node, err
}
