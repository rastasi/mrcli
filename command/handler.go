package command

import (
	"flag"
	"mrcli/lib/config"
	"mrcli/lib/logger"
)

func setFlags(commandProperties []commandProperty) {
	for i, _ := range commandProperties {
		flag.StringVar(&commandProperties[i].Value, commandProperties[i].Id, "", commandProperties[i].Description)
	}
}

func handle(commandProperties []commandProperty) bool {
	for _, c := range commandProperties {
		if c.Value != "" {
			if c.checkConfigFile && !config.CheckConfigFile(".") {
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
		logger.Log("", "%s: %s\n", f.Name, f.Usage)
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
