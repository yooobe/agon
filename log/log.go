package log

import (
	"io"
	"fmt"
	"os"
	"github.com/biezhi/agon/color"
	"time"
)

type Level int

// log level
const (
	LEVEL_DEBUG  = iota
	LEVEL_TRACE
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
)

// print info log
func Info(format string, v...interface{}) {
	Log(os.Stdout, LEVEL_INFO, fmt.Sprintf(format, v...))
}

// print trace log
func Trace(format string, v...interface{}) {
	Log(os.Stdout, LEVEL_TRACE, fmt.Sprintf(format, v...))
}

// print debug log
func Debug(format string, v...interface{}) {
	Log(os.Stdout, LEVEL_DEBUG, fmt.Sprintf(format, v...))
}

// print warning log
func Warn(format string, v...interface{}) {
	Log(os.Stdout, LEVEL_WARN, fmt.Sprintf(format, v...))
}

// print error log
func Error(format string, v...interface{}) {
	Log(os.Stdout, LEVEL_ERROR, fmt.Sprintf(format, v...))
}

// get log level text
func LevelText(level Level) string {
	switch {
	case level == LEVEL_INFO:
		return "INFO"
	case level == LEVEL_DEBUG:
		return "DEBUG"
	case level == LEVEL_TRACE:
		return "TRACE"
	case level == LEVEL_WARN:
		return "WARN"
	case level == LEVEL_ERROR:
		return "ERROR"
	default:
		return ""
	}
}

// get log level color
func Color(str string, level Level) string {
	switch {
	case level == LEVEL_INFO:
		return color.Color(color.Green, color.Normal, str)
	case level == LEVEL_DEBUG:
		return color.Color(color.LightGray, color.Normal, str)
	case level == LEVEL_TRACE:
		return color.Color(color.LightPurple, color.Normal, str)
	case level == LEVEL_WARN:
		return color.Color(color.Yellow, color.Normal, str)
	case level == LEVEL_ERROR:
		return color.Color(color.Red, color.Normal, str)
	default:
		return color.Color(color.Cyan, color.Normal, str)
	}
}

// return time log
func TimeLog() string {
	return time.Now().Format("01/02 15:04:05.000")
}

// print log
func Log(out io.Writer, level Level, txt string) {
	prefix := fmt.Sprintf("[%s] %s\t=> ", TimeLog(), LevelText(level))
	out.Write([]byte(Color(prefix, level)))
	out.Write([]byte(color.Color(color.LightGray, color.Normal, txt) + "\n"))
}
