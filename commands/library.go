package commands

import "mrcli/lib"

func Library(name string) {
	lib.LogSuccess(name, "Initialize application\n", []any{})

	lib.CreateDir("libs/" + name)
	lib.CreateGoWorkFile(name, "./libs/"+name)
}
