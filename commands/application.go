package commands

import "mrcli/lib"

func Application(name string) {
	lib.LogSuccess(name, "Initialize application\n", []any{})
	lib.CreateDir("apps/" + name)
	lib.CreateGoWorkFile(name, "./apps/"+name)
}
