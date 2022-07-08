package commands

import (
	"mrcli/lib/config"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
)

func checkBaseDirs(name string) {
	logger.Log(name, "Base directories")
	filemanager.FileExistsDisplayBulk([]string{
		name + "/project.yaml",
		name + "/apps",
		name + "/libs",
		name + "/dists",
	})
}

func checkConfigFileNilDisplay(propertyName, property string) {
	if len(property) != 0 {
		logger.LogSuccess(propertyName, "Property is filled")
	} else {
		logger.LogFail(propertyName, "Property is missing")
	}
}

func checkConfigFileFormat(name string) {
	logger.Log(name, "Config file properties\n")
	if config.CheckConfigFile(name) {
		config := config.LoadConfigFile(name)
		checkConfigFileNilDisplay("Name", config.Name)
		checkConfigFileNilDisplay("Type", config.Type)
		checkConfigFileNilDisplay("Version", config.Version)
	} else {
		logger.Log(name, "Config file not found or incorrect")
	}
}

func Check(name string) {
	logger.Log(name, "Check %s\n", name)
	checkBaseDirs(name)
	checkConfigFileFormat(name)
}
