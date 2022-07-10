package command

import (
	"mrcli/lib/config"
	"mrcli/lib/exec"
	"mrcli/lib/logger"
	"mrcli/lib/metadata"
	"mrcli/lib/structure"
)

func Initialize(projectName string) {
	logger.LogSuccess(projectName, "Initialize monorepository\n")
	info := metadata.GetMetadata()
	structure.BuildFromStructure(info.Project.Structure, projectName)
	config.CreateConfigFile(projectName)
	exec.CreateGoModFile(projectName)
	exec.CreateGitRepo(projectName)
}
