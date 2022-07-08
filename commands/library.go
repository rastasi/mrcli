package commands

import (
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
)

func Library(name string) {
	logger.LogSuccess(name, "Initialize application\n", []any{})
	filemanager.CreateDir("libs/" + name)
	exec.CreateGoWorkFile(name, "./libs/"+name)
}
