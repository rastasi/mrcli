package commands

import (
	"mrcli/lib/config"
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
)

func Initialize(name string) {
	logger.LogSuccess(name, "Initialize monorepository\n", []any{})
	filemanager.CreateDirBulk([]string{
		"./" + name,
		"./" + name + "/apps",
		"./" + name + "/libs",
		"./" + name + "/dists",
	})
	config.CreateConfigFile(name)
	exec.CreateGoModFile(name)
	exec.CreateGitRepo((name))
}
