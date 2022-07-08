package commands

import (
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
)

func Application(name string) {
	logger.LogSuccess(name, "Initialize application\n", []any{})
	filemanager.CreateDirBulk([]string{
		"apps/" + name,
		"dists/" + name,
	})
	filemanager.CreateFile("./dists/" + name + "/.keep")
	exec.CreateGoWorkFile(name, "./apps/"+name)
}
