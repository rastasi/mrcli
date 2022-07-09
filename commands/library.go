package commands

import (
	"mrcli/lib/config"
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
)

func Library(name string) {
	logger.LogSuccess(name, "Initialize application\n")
	info := config.GetMetadata()
	filemanager.CreateDirBulk(info.Library.BaseDirectoryStructure, name)
	filemanager.CreateFileBulk(info.Library.BaseFiles, name)
	exec.CreateGoWorkFile(name, "./libs/"+name)
}
