package commands

import (
	"mrcli/lib/config"
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
)

func Application(projectName string) {
	logger.LogSuccess(projectName, "Initialize application\n")
	info := config.GetMetadata()
	filemanager.CreateDirBulk(info.Application.BaseDirectoryStructure, projectName)
	filemanager.CreateFileBulk(info.Application.BaseFiles, projectName)
	exec.CreateGoWorkFile(projectName, "./apps/"+projectName)
}
