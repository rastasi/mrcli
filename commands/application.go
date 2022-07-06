package commands

import (
	"mrcli/utils"
)

func Application(name string) {
	utils.LogSuccess(name, "Initialize application\n", []any{})
	utils.CreateDir("apps/" + name)
	utils.CreateGoWorkFile(name, "./apps/"+name)
}
