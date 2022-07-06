package lib

func CreateGoModFile(name string) {
	ExecuteCommand(name, name, "go", []string{"mod", "init", name})
}

func CreateGoWorkFile(name string, subdir string) {
	ExecuteCommand(name, subdir, "go", []string{"work", "init"})
}
