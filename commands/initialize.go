package commands

import (
	"mrcli/utils"
)

func Initialize(name string) {
	utils.LogSuccess(name, "Initialize monorepository\n", []any{})
	utils.CreateDir("./" + name)
	utils.CreateDir("./" + name + "/apps")
	utils.CreateDir("./" + name + "/libs")
	utils.CreateConfigFile(name)
	utils.CreateGoModFile(name)
}
