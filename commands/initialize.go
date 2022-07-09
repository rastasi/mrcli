package commands

import (
	"mrcli/lib/config"
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
)

func Initialize(projectName string) {
	logger.LogSuccess(projectName, "Initialize monorepository\n")
	info := config.GetMetadata()
	filemanager.CreateDirBulk(info.Project.BaseDirectoryStructure, projectName)
	filemanager.CreateFileBulk(info.Project.BaseFiles, projectName)
	config.CreateConfigFile(projectName)
	exec.CreateGoModFile(projectName)
	exec.CreateGitRepo(projectName)
}
