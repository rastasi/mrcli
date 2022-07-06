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

func CreateConfigfile(name string) {

	configData, err := yaml.Marshal(Configfile{
		Type:    "mrcli",
		Version: "1.0",
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
