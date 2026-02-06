package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL

	// ANSI color codes

	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorGray   = "\033[90m"

	// Bold colors
	ColorBoldRed    = "\033[1;31m"
	ColorBoldGreen  = "\033[1;32m"
	ColorBoldYellow = "\033[1;33m"
	ColorBoldBlue   = "\033[1;34m"

	// Background colors
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
)

type Dot struct {
	level      LogLevel
	showSource bool
}

var defaultDot = &Dot{
	level:      INFO,
	showSource: false,
}

func NewDot(level LogLevel) *Dot {
	return &Dot{
		level:      level,
		showSource: true,
	}
}

func (d *Dot) getSource() string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return ""
	}

	parts := strings.Split(file, "/")
	filename := parts[len(parts)-1]

	return fmt.Sprintf("%s:%d", filename, line)
}

func (d *Dot) log(level LogLevel, levelName, color, message string, err error) {
	if level < d.level {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	source := ""

	if d.showSource {
		source = fmt.Sprintf(" %s[%s]%s", ColorGray, d.getSource(), ColorReset)
	}

	fmt.Printf(
		"%s%s%s %s%-8s%s %s%s%s\n",
		ColorGray,
		timestamp,
		ColorReset,
		color,
		"["+levelName+"]",
		ColorReset,
		color,
		message,
		ColorReset,
	)

	if err != nil {
		fmt.Printf("%s Error: %v%s\n", color, err, ColorReset)
	}

	if d.showSource && source != "" {
		fmt.Printf("%s Source:%s\n", ColorGray, source)
	}

	if level == FATAL {
		os.Exit(1)
	}
}

func (d *Dot) Debug(message string) {
	d.log(DEBUG, "DEBUG", ColorCyan, message, nil)
}

func (d *Dot) Info(message string) {
	d.log(INFO, "INFO", ColorBlue, message, nil)
}

func (d *Dot) Success(message string) {
	d.log(INFO, "SUCCESS", ColorGreen, message, nil)
}

func (d *Dot) Warning(message string) {
	d.log(WARNING, "WARNING", ColorYellow, message, nil)
}

func (d *Dot) Error(message string) {
	d.log(ERROR, "ERROR", ColorRed, message, nil)
}

func (d *Dot) Fatal(message string) {
	d.log(FATAL, "FATAL", ColorRed, message, nil)
}

func Debug(message string) {
	defaultDot.Debug(message)
}

func Info(message string) {
	defaultDot.Info(message)
}

func Success(message string) {
	defaultDot.Success(message)
}

func Warning(message string) {
	defaultDot.Warning(message)
}

func Error(message string) {
	defaultDot.Error(message)
}

func Fatal(message string) {
	defaultDot.Fatal(message)
}

func SetLogLevel(level LogLevel) {
	defaultDot.level = level
}
