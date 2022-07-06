package commands

import (
	"fmt"
	"mrcli/utils"
)

func Library(name string) {
	fmt.Printf("Initialize application: %s\n", name)
	utils.CreateDir("libs/" + name)
	utils.CreateGoWorkFile(name, "./libs/"+name)
}
