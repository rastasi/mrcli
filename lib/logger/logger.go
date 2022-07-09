package logger

import (
	"fmt"

	"github.com/fatih/color"
)

func formatLog(prefix string, message string) string {
	if len(prefix) == 0 {
		return fmt.Sprintf("%s\n", message)
	} else {
		return fmt.Sprintf("[%s] %s\n", prefix, message)
	}
}

func Log(prefix string, message string, args ...any) {
	formatted := formatLog(prefix, message)
	fmt.Printf(formatted, args...)
}

func LogSuccess(prefix string, message string, args ...any) {
	formatted := formatLog(prefix, message)
	color.New(color.FgGreen).Printf(formatted, args...)
}

func LogFail(prefix string, message string, args ...any) {
	formatted := formatLog(prefix, message)
	color.New(color.FgRed).Printf(formatted, args...)
}
