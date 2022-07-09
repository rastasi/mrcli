package commands

import (
	"mrcli/lib/config"
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
)

func Application(name string) {
	logger.LogSuccess(name, "Initialize application\n")
	info := config.GetMetadata()
	filemanager.CreateDirBulk(info.Application.BaseDirectoryStructure, name)
	filemanager.CreateFileBulk(info.Application.BaseFiles, name)
	filemanager.CreateFile("./dists/" + name + "/.keep")
	exec.CreateGoWorkFile(name, "./apps/"+name)
}
