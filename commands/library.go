package commands

import (
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
	"mrcli/lib/metadata"
)

func Library(libraryName string) {
	logger.LogSuccess(libraryName, "Initialize application\n")
	info := metadata.GetMetadata()
	filemanager.BuildFromStructure(info.Library.Structure, libraryName)
	exec.CreateGoWorkFile(libraryName, "./libs/"+libraryName)
}
