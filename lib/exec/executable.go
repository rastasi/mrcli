package exec

func CreateGoModFile(name string) {
	ExecuteCommand(name, name, "go", []string{"mod", "init", name})
}

func CreateGoWorkFile(name string, subdir string) {
	ExecuteCommand(name, subdir, "go", []string{"work", "init"})
}

func CreateGitRepo(name string) {
	ExecuteCommand(name, name, "git", []string{"init"})
}
