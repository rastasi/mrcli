package main

import (
	"flag"
	"mrcli/commands"
)

func main() {
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
}
