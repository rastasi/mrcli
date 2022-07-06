package lib

import (
	"os/exec"
	"strings"
)

func ExecuteCommand(name string, subdir string, command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Dir = subdir
	err := cmd.Run()
	Log(name, "Execute command: \"%s %s\" in \"%s\"\n", []any{command, strings.Join(args, " "), subdir})
	if err != nil {
		LogFail(name, err.Error()+"\n", []any{})
	}
}
