package commands

import (
	"flag"
	"mrcli/lib"
)

func setFlags(commandProperties []commandProperty) {
	for i, _ := range commandProperties {
		flag.StringVar(&commandProperties[i].Value, commandProperties[i].Id, "", commandProperties[i].Description)
	}
}

func handle(commandProperties []commandProperty) bool {
	for _, c := range commandProperties {
		if c.Value != "" {
			if c.checkConfigFile && !lib.CheckConfigFile(".") {
				panic("Config file not found or incorrect")
			}
			c.Handler(c.Value)
			return true
		}
	}
	return false
}

func help() {
	flag.VisitAll(func(f *flag.Flag) {
		lib.Log("", "%s: %s\n", []any{f.Name, f.Usage})
	})
}

func CommandHandler() {
	commandProperties := GetCommandProperties()
	setFlags(commandProperties)
	flag.Parse()
	if !handle(commandProperties) {
		help()
	}
}
