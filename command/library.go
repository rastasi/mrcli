package command

import (
	"mrcli/lib/exec"
	"mrcli/lib/logger"
	"mrcli/lib/metadata"
	"mrcli/lib/structure"
)

func Library(libraryName string) {
	logger.LogSuccess(libraryName, "Initialize library\n")
	info := metadata.GetMetadata()
	structure.BuildFromStructure(info.Library.Structure, libraryName)
	exec.CreateGoWorkFile(libraryName, "./libs/"+libraryName)
}
