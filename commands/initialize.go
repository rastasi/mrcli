package commands

import "mrcli/lib"

func Initialize(name string) {
	lib.LogSuccess(name, "Initialize monorepository\n", []any{})
	lib.CreateDir("./" + name)
	lib.CreateDir("./" + name + "/apps")
	lib.CreateDir("./" + name + "/libs")
	lib.CreateConfigFile(name)
	lib.CreateGoModFile(name)
}
