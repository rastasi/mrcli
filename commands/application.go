package commands

import "mrcli/lib"

func Application(name string) {
	lib.LogSuccess(name, "Initialize application\n", []any{})
	lib.CreateDirBulk([]string{
		"apps/" + name,
		"dists/" + name,
	})
	lib.CreateFile("./dists/" + name + "/.keep")
	lib.CreateGoWorkFile(name, "./apps/"+name)
}
