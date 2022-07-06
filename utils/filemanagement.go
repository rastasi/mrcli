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
		fmt.Println(err)
	}
}

func CreateDir(name string) {
	if !exists(name) {
		permission := 0766
		fmt.Println("Create directory: " + name)
		os.Mkdir(name, fs.FileMode(permission))
	} else {
		fmt.Println("Directory already exists: " + name)
	}
}

func CreateGoModFile(name string) {
	executeCommand(name, name, "go", []string{"mod", "init", name})
}

func CreateGoWorkFile(name string, subdir string) {
	executeCommand(name, subdir, "go", []string{"work", "init"})
}
