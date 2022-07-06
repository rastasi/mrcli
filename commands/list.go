package commands

type commandProperty struct {
	Id           string
	Description  string
	Handler      func(name string)
	Value        string
	InProjectDir bool
}

func GetCommandProperties() []commandProperty {
	return []commandProperty{
		{
			Id:          "init",
			Description: "Initialize monorepository",
			Handler:     Initialize,
		},
		{
			Id:           "app",
			Description:  "Create application",
			Handler:      Application,
			InProjectDir: true,
		},
		{
			Id:           "lib",
			Description:  "Create library",
			Handler:      Library,
			InProjectDir: true,
		},
		{
			Id:           "build",
			Description:  "Build application",
			Handler:      Build,
			InProjectDir: true,
		},
	}
}
