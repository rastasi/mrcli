package commands

import (
	"mrcli/layout"
	"mrcli/utils"
)

func Library(name string) {
	layout.LogSuccess(name, "Initialize application\n", []any{})

	utils.CreateDir("libs/" + name)
	utils.CreateGoWorkFile(name, "./libs/"+name)
}
