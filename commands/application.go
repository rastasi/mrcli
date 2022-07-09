package commands

import (
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
	"mrcli/lib/metadata"
)

func Application(appName string) {
	logger.LogSuccess(appName, "Initialize application\n")
	info := metadata.GetMetadata()
	filemanager.BuildFromStructure(info.Application.Structure, appName)
	exec.CreateGoWorkFile(appName, "./apps/"+appName)
}
