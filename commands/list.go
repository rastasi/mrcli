package commands

type commandProperty struct {
	Id              string
	Description     string
	Handler         func(name string)
	Value           string
	checkConfigFile bool
}

func GetCommandProperties() []commandProperty {
	return []commandProperty{
		{
			Id:          "init",
			Description: "Initialize monorepository",
			Handler:     Initialize,
		},
		{
			Id:              "app",
			Description:     "Create application",
			Handler:         Application,
			checkConfigFile: true,
		},
		{
			Id:              "lib",
			Description:     "Create library",
			Handler:         Library,
			checkConfigFile: true,
		},
		{
			Id:              "build",
			Description:     "Build application",
			Handler:         Build,
			checkConfigFile: true,
		},
		{
			Id:              "check",
			Description:     "Check monorepository structure",
			Handler:         Check,
			checkConfigFile: false,
		},
	}
}
