package lib

import (
	"fmt"

	"github.com/fatih/color"
)

func formatLog(name string, message string) string {
	var formatted string
	if len(name) == 0 {
		formatted = fmt.Sprintf("%s", message)
	} else {
		formatted = fmt.Sprintf("[%s] %s", name, message)
	}
	return formatted
}

func Log(name string, message string, args []any) {
	formatted := formatLog(name, message)
	fmt.Printf(formatted, args...)
}

func LogSuccess(name string, message string, args []any) {
	formatted := formatLog(name, message)
	color.New(color.FgGreen).Printf(formatted, args...)
}

func LogFail(name string, message string, args []any) {
	formatted := formatLog(name, message)
	color.New(color.FgRed).Printf(formatted, args...)
}
