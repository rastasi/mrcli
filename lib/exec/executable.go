package exec

func CreateGoModFile(projectName string) {
	ExecuteCommand(projectName, projectName, "go", []string{"mod", "init", projectName})
}

func CreateGoWorkFile(projectName string, subdir string) {
	ExecuteCommand(projectName, subdir, "go", []string{"work", "init"})
}

func CreateGitRepo(projectName string) {
	ExecuteCommand(projectName, projectName, "git", []string{"init"})
}
