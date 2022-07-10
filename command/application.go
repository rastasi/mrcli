package command

import (
	"mrcli/lib/exec"
	"mrcli/lib/logger"
	"mrcli/lib/metadata"
	"mrcli/lib/structure"
)

func Application(appName string) {
	logger.LogSuccess(appName, "Initialize application\n")
	info := metadata.GetMetadata()
	structure.BuildFromStructure(info.Application.Structure, appName)
	exec.CreateGoWorkFile(appName, "./apps/"+appName)
}
