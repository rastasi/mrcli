package commands

import (
	"mrcli/lib/config"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
)

func checkBaseDirs(projectName string) {
	logger.Log(projectName, "Base directories")
	info := config.GetMetadata()
	filemanager.FileExistsDisplayBulk(info.Project.BaseDirectoryStructure, "")
}

func checkConfigFileNilDisplay(propertyName, property string) {
	if len(property) != 0 {
		logger.LogSuccess(propertyName, "Property is filled")
	} else {
		logger.LogFail(propertyName, "Property is missing")
	}
}

func checkConfigFileFormat(projectName string) {
	logger.Log(projectName, "Config file properties\n")
	if config.CheckConfigFile(projectName) {
		config := config.LoadConfigFile(projectName)
		checkConfigFileNilDisplay("Name", config.Name)
		checkConfigFileNilDisplay("Type", config.Type)
		checkConfigFileNilDisplay("Version", config.Version)
	} else {
		logger.Log(projectName, "Config file not found or incorrect")
	}
}

func Check(projectName string) {
	logger.Log(projectName, "Check %s\n", projectName)
	checkBaseDirs(projectName)
	checkConfigFileFormat(projectName)
}
