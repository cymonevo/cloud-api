package log

import (
	"github.com/fatih/color"
	"time"
)

const (
	tagline = "[%s] "
	errline = ", Err: %v"
	newline = "\n"
)

var (
	timeline = time.Now().Format("2006-01-02 15:04:05 ")

	white   = color.New(color.FgWhite)
	green   = color.New(color.FgGreen)
	blue    = color.New(color.FgBlue)
	yellow  = color.New(color.FgYellow)
	red     = color.New(color.FgRed)
	redBold = color.New(color.FgRed, color.Bold)

	print = white.PrintlnFunc()
	info  = green.PrintlnFunc()
	debug = blue.PrintlnFunc()
	warn  = yellow.PrintlnFunc()
	error = red.PrintlnFunc()
	fatal = redBold.PrintlnFunc()

	printf = white.PrintfFunc()
	infof  = green.PrintfFunc()
	debugf = blue.PrintfFunc()
	warnf  = yellow.PrintfFunc()
	errorf = red.PrintfFunc()
	fatalf = redBold.PrintfFunc()
)
