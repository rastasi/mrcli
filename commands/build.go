package commands

import "mrcli/lib"

func Build(name string) {
	lib.ExecuteCommand(name, "./apps/"+name, "go", []string{"build", "."})
}
