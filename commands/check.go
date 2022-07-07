package commands

import "mrcli/lib"

func checkBaseDirs(name string) {
	lib.Log(name, "Base directories\n", []any{})
	lib.FileExistsDisplayBulk([]string{
		name + "/project.yaml",
		name + "/apps",
		name + "/libs",
		name + "/dists",
	})
}

func checkConfigFileNilDisplay(propertyName, property string) {
	if len(property) != 0 {
		lib.LogSuccess(propertyName, "Property is filled\n", []any{})
	} else {
		lib.LogFail(propertyName, "Property is missing\n", []any{})
	}
}

func checkConfigFileFormat(name string) {
	lib.Log(name, "Config file properties\n", []any{})
	if lib.CheckConfigFile(name) {
		config := lib.LoadConfigFile(name)
		checkConfigFileNilDisplay("Name", config.Name)
		checkConfigFileNilDisplay("Type", config.Type)
		checkConfigFileNilDisplay("Version", config.Version)
	} else {
		lib.Log(name, "Config file not found or incorrect", []any{})
	}
}

func Check(name string) {
	lib.Log(name, "Check %s\n", []any{name})
	checkBaseDirs(name)
	checkConfigFileFormat(name)
}
