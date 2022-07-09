package commands

import "mrcli/lib/exec"

func Build(appName string) {
	exec.ExecuteCommand(appName, "./apps/"+appName, "go", []string{"build", "."})
}
