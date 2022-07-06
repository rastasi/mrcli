package main

import (
	"mrcli/commandhandler"
	"mrcli/layout"
)

func main() {
	layout.Header()
	commandhandler.CommandHandler()
}
