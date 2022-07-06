package commands

import (
	"fmt"
	"mrcli/utils"
)

func Application(name string) {
	fmt.Printf("Initialize application: %s\n", name)
	utils.CreateDir("apps/" + name)
	utils.CreateGoWorkFile(name, "./apps/"+name)
}
