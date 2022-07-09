package commands

import (
	"mrcli/lib/config"
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
)

func Library(libraryName string) {
	logger.LogSuccess(libraryName, "Initialize application\n")
	info := config.GetMetadata()
	filemanager.CreateDirBulk(info.Library.BaseDirectoryStructure, libraryName)
	filemanager.CreateFileBulk(info.Library.BaseFiles, libraryName)
	exec.CreateGoWorkFile(libraryName, "./libs/"+libraryName)
}
