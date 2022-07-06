package commands

import (
	"fmt"
	"mrcli/utils"
)

func Initialize(name string) {
	fmt.Println("Initialize monorepository")
	utils.CreateDir("./" + name)
	utils.CreateDir("./" + name + "/apps")
	utils.CreateDir("./" + name + "/libs")
	utils.CreateConfigFile(name)
	utils.CreateGoModFile(name)
}
