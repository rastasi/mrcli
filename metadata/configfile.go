package metadata

import (
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
	configData, _ := yaml.Marshal(Configfile{
		Type:    metadata.ProjectName,
		Version: metadata.Version,
		Name:    name,
	})

	path := name + "/project.yaml"
	ioutil.WriteFile(path, configData, 0644)
}
