package commands

import "mrcli/lib"

func Initialize(name string) {
	lib.LogSuccess(name, "Initialize monorepository\n", []any{})
	lib.CreateDirBulk([]string{
		"./" + name,
		"./" + name + "/apps",
		"./" + name + "/libs",
		"./" + name + "/dists",
	})
	lib.CreateConfigFile(name)
	lib.CreateGoModFile(name)
	lib.CreateGitRepo((name))
}
