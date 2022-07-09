package exec

import (
	"mrcli/lib/logger"
	"os/exec"
	"strings"
)

func ExecuteCommand(name string, subdir string, command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Dir = subdir
	err := cmd.Run()
	logger.Log(name, "Execute command: \"%s %s\" in \"%s\"\n", command, strings.Join(args, " "), subdir)
	if err != nil {
		logger.LogFail(name, err.Error()+"\n")
	}
}
