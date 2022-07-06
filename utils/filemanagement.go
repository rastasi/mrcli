package utils

import (
	"io/fs"
	"mrcli/layout"
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

func ExecuteCommand(name string, subdir string, command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Dir = subdir
	err := cmd.Run()
	layout.Log(name, "Execute command: \"%s %s\" in \"%s\"\n", []any{command, strings.Join(args, " "), subdir})
	if err != nil {
		layout.LogFail(name, err.Error()+"\n", []any{})
	}
}

func CreateDir(name string) {
	if !exists(name) {
		permission := 0766
		layout.LogSuccess(name, "Create directory\n", []any{})
		os.Mkdir(name, fs.FileMode(permission))
	} else {
		layout.LogFail(name, "Directory already exists\n", []any{})
	}
}
