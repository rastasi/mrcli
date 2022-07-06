package utils

func CreateGoModFile(name string) {
	executeCommand(name, name, "go", []string{"mod", "init", name})
}

func CreateGoWorkFile(name string, subdir string) {
	executeCommand(name, subdir, "go", []string{"work", "init"})
}
