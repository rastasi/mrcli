package config

import (
	"io/ioutil"
	"mrcli/lib/filemanager"

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
		Type:    metadata.Name,
		Version: metadata.Version,
		Name:    name,
	})

	path := name + "/project.yaml"
	ioutil.WriteFile(path, configData, 0644)
}

func LoadConfigFile(name string) Configfile {
	fileData, _ := ioutil.ReadFile(name + "/project.yaml")
	var config Configfile
	yaml.Unmarshal(fileData, &config)
	return config
}

func CheckConfigFile(name string) bool {
	if filemanager.FileExists(name + "/project.yaml") {
		config := LoadConfigFile(name)
		metadata := GetMetadata()
		if config.Type == metadata.Name {
			return true
		}
		return false
	}
	return false
}
