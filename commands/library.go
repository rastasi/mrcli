package commands

import (
	"mrcli/utils"
)

func Library(name string) {
	utils.LogSuccess(name, "Initialize application\n", []any{})

	utils.CreateDir("libs/" + name)
	utils.CreateGoWorkFile(name, "./libs/"+name)
}
