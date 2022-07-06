package utils

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func executeCommand(name string, subdir string, command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Dir = subdir
	err := cmd.Run()
	fmt.Printf("Execute command: \"%s %s\" in \"%s\"\n", command, strings.Join(args, " "), subdir)
	if err != nil {
		LogFail(name, err.Error()+"\n", []any{})
	}
}

func CreateDir(name string) {
	if !exists(name) {
		permission := 0766
		LogSuccess(name, "Create directory\n", []any{})
		os.Mkdir(name, fs.FileMode(permission))
	} else {
		LogFail(name, "Directory already exists\n", []any{})
	}
}

func CreateGoModFile(name string) {
	executeCommand(name, name, "go", []string{"mod", "init", name})
}

func CreateGoWorkFile(name string, subdir string) {
	executeCommand(name, subdir, "go", []string{"work", "init"})
}
