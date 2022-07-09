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

func CreateConfigFile(projectName string) {
	info := GetMetadata()
	configData, _ := yaml.Marshal(Configfile{
		Type:    info.Name,
		Version: info.Version,
		Name:    projectName,
	})

	path := projectName + "/project.yaml"
	ioutil.WriteFile(path, configData, 0644)
}

func LoadConfigFile(projectName string) Configfile {
	fileData, _ := ioutil.ReadFile(projectName + "/project.yaml")
	var config Configfile
	yaml.Unmarshal(fileData, &config)
	return config
}

func CheckConfigFile(projectName string) bool {
	if filemanager.FileExists(projectName + "/project.yaml") {
		config := LoadConfigFile(projectName)
		metadata := GetMetadata()
		if config.Type == metadata.Name {
			return true
		}
		return false
	}
	return false
}
