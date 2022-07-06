package main

import (
	"mrcli/commands"
	"mrcli/lib"
)

func main() {
	lib.Header()
	commands.CommandHandler()
}
