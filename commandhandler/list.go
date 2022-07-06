package commandhandler

import "mrcli/commands"

type commandProperty struct {
	Id          string
	Description string
	Handler     func(name string)
	Value       string
}

func GetCommandProperties() []commandProperty {
	return []commandProperty{
		{
			Id:          "init",
			Description: "Initialize monorepository",
			Handler:     commands.Initialize,
		},
		{
			Id:          "app",
			Description: "Initialize application",
			Handler:     commands.Application,
		},
		{
			Id:          "lib",
			Description: "Initialize library",
			Handler:     commands.Library,
		},
	}
}
