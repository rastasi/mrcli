package commands

import "mrcli/lib/exec"

func Build(name string) {
	exec.ExecuteCommand(name, "./apps/"+name, "go", []string{"build", "."})
}
