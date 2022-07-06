package commands

import (
	"mrcli/layout"
	"mrcli/metadata"
	"mrcli/utils"
)

func Initialize(name string) {
	layout.LogSuccess(name, "Initialize monorepository\n", []any{})
	utils.CreateDir("./" + name)
	utils.CreateDir("./" + name + "/apps")
	utils.CreateDir("./" + name + "/libs")
	metadata.CreateConfigFile(name)
	utils.CreateGoModFile(name)
}
