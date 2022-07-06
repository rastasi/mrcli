package commands

import (
	"mrcli/layout"
	"mrcli/utils"
)

func Application(name string) {
	layout.LogSuccess(name, "Initialize application\n", []any{})
	utils.CreateDir("apps/" + name)
	utils.CreateGoWorkFile(name, "./apps/"+name)
}
