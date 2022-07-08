package logger

import (
	"fmt"

	"github.com/fatih/color"
)

func formatLog(name string, message string) string {
	if len(name) == 0 {
		return fmt.Sprintf("%s\n", message)
	} else {
		return fmt.Sprintf("[%s] %s\n", name, message)
	}
}

func Log(name string, message string, args ...any) {
	formatted := formatLog(name, message)
	fmt.Printf(formatted, args...)
}

func LogSuccess(name string, message string, args ...any) {
	formatted := formatLog(name, message)
	color.New(color.FgGreen).Printf(formatted, args...)
}

func LogFail(name string, message string, args ...any) {
	formatted := formatLog(name, message)
	color.New(color.FgRed).Printf(formatted, args...)
}
