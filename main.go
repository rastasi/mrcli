package main

import (
	"flag"
	"fmt"
	"mrcli/commands"
	"mrcli/utils"
)

func main() {
	metadata := utils.GetMetadata()
	fmt.Printf("%s %s\n\n", metadata.ProjectName, metadata.Version)

	var mrCmd string
	flag.StringVar(&mrCmd, "init", "", "Initialize monorepository")

	var appCmd string
	flag.StringVar(&appCmd, "app", "", "Initialize application")

	var libCmd string
	flag.StringVar(&appCmd, "lib", "", "Initialize shared library")

	flag.Parse()

	if mrCmd != "" {
		commands.Initialize(mrCmd)
		return
	}

	if appCmd != "" {
		commands.Application(appCmd)
		return
	}

	if libCmd != "" {
		commands.Library(libCmd)
		return
	}

	commands.Help()
}
