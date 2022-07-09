package commands

import (
	"mrcli/lib/config"
	"mrcli/lib/exec"
	"mrcli/lib/filemanager"
	"mrcli/lib/logger"
	"mrcli/lib/metadata"
)

func Initialize(projectName string) {
	logger.LogSuccess(projectName, "Initialize monorepository\n")
	info := metadata.GetMetadata()
	filemanager.BuildFromStructure(info.Project.Structure, projectName)
	config.CreateConfigFile(projectName)
	exec.CreateGoModFile(projectName)
	exec.CreateGitRepo(projectName)
}
