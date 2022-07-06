package commandhandler

import (
	"flag"
	"fmt"
	"mrcli/commands"
)

func CommandHandler() {
	commandProperties := GetCommandProperties()

	for i, _ := range commandProperties {
		fmt.Println(commandProperties[i].Id)
		flag.StringVar(&commandProperties[i].Value, commandProperties[i].Id, "", commandProperties[i].Description)
	}

	flag.Parse()

	for _, c := range commandProperties {
		fmt.Println(c.Value)
		if c.Value != "" {
			c.Handler(c.Value)
			return
		}
	}

	commands.Help()
}
