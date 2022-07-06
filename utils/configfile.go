package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configfile struct {
	Type    string `yaml:"type"`
	Version string `yaml:"version"`
	Name    string `yaml:"name"`
}

func CreateConfigFile(name string) {
	metadata := GetMetadata()
	configData, err := yaml.Marshal(Configfile{
		Type:    metadata.ProjectName,
		Version: metadata.Version,
		Name:    name,
	})

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	path := name + "/project.yaml"

	err = ioutil.WriteFile(path, configData, 0644)

	if err != nil {
		panic("Unable to write data into the file")
	}
}
