package commands

import (
	"mrcli/lib/config"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
	"mrcli/lib/metadata"
)

func getBaseDirPatterns() []string {
	info := metadata.GetMetadata()
	var patterns []string
	for _, entity := range info.Project.Structure {
		patterns = append(patterns, entity.Pattern)
	}
	return patterns
}

func checkBaseDirs(projectName string) {
	logger.Log(projectName, "Base directories")
	patterns := getBaseDirPatterns()
	filemanager.FileExistsDisplayBulk(patterns, "")
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
