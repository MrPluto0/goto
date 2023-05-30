package log

import "github.com/fatih/color"

func Info(format string, a ...interface{}) {
	color.Cyan(format, a...)
}
