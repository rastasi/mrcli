package commands

import (
	"mrcli/lib/config"
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
)

func Initialize(name string) {
	logger.LogSuccess(name, "Initialize monorepository\n")
	info := config.GetMetadata()
	filemanager.CreateDirBulk(info.Project.BaseDirectoryStructure, name)
	filemanager.CreateFileBulk(info.Project.BaseFiles, name)
	config.CreateConfigFile(name)
	exec.CreateGoModFile(name)
	exec.CreateGitRepo(name)
}
