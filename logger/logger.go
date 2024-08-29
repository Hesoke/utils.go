// JUST SIMPLE ASYNC LOGGER WITH BASIC OPTIONS
// THERE WOULD BE NO PATTERN SUPPORT, FAIL-SAFE MECHANISMS ETC.
package logger

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"time"

	"github.com/bigelle/utils.go/ensure"
)

// Logging level preseneted as int, where 0 is DEBUG, and 4 is FATAL
type LoggingLevel int

const (
	DEBUG LoggingLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

func (l LoggingLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return fmt.Sprintf("LEVEL:%d", l)
	}
}

type color int

const (
	ColorReset = "\033[0m"
	ColorDebug = "\033[36m"
	ColorInfo  = "\033[97m"
	ColorWarn  = "\033[33m"
	ColorError = "\033[31m"
	ColorFatal = "\033[1;31m"
)

func (c color) String() string {
	switch c {
	case color((DEBUG)):
		return ColorDebug
	case color(INFO):
		return ColorInfo
	case color(WARN):
		return ColorWarn
	case color(ERROR):
		return ColorError
	case color(FATAL):
		return ColorFatal
	default:
		return ColorReset
	}
}

type logger struct {
	Level LoggingLevel
	// IMPORTANT: if you want to write logs to file, open your file in appending mode
	Writer io.Writer
	chLog  chan string
}

// global var used to store configuration without creating local struct instance every time
// across the project
var global logger

// functional option that is used while initializing logger using NewLogger()
type Opt func(*logger)

// Minimal allowed loggin level
func WithLevel(level int) Opt {
	return func(l *logger) {
		l.Level = LoggingLevel(level)
	}
}

// Writer that would be used to print logs to.
// IMPORTANT: if you want to write logs to file, open your file in appending mode
func WithWriter(w io.Writer) Opt {
	return func(l *logger) {
		if w != nil {
			l.Writer = bufio.NewWriter(w)
		}
	}
}

// Initializes logger globally with options.
// IMPORTANT: if you want to write logs to file, open your file in appending mode
func NewLogger(opts ...Opt) {
	global = logger{
		Level:  INFO,
		Writer: os.Stdout,
	}

	for _, opt := range opts {
		opt(&global)
	}
	global.chLog = make(chan string)
	printer()
}

// Log is loggin msg to io.Writer with logging lvl and source of message src in format:
//
// [HH:MM:SS] [lvl] src: msg
//
// IMPORTANT: if you want to write logs to file, open your file in appending mode
func Log[L int | LoggingLevel](lvl L, src string, msg string) {
	if !ensure.NotNil(global) {
		return
	}
	if int(lvl) < int(global.Level) {
		return
	}
	global.chLog <- fill(lvl, src, msg)
}

// Debug is logging msg from src to io.Writer with DEBUG logging level
//
// [HH:MM:SS] [DEBUG] src: msg
//
// IMPORTANT: if you want to write logs to file, open your file in appending mode
func Debug(src string, msg string) {
	if !ensure.NotNil(global) {
		return
	}
	if 0 < global.Level {
		return
	}
	global.chLog <- fill(0, src, msg)
}

// Info is logging msg from src to io.Writer with INFO logging level
//
// [HH:MM:SS] [INFO] src: msg
//
// IMPORTANT: if you want to write logs to file, open your file in appending mode
func Info(src string, msg string) {
	if !ensure.NotNil(global) {
		return
	}
	if 1 < global.Level {
		return
	}
	global.chLog <- fill(1, src, msg)
}

// Warn is logging msg from src to io.Writer with WARN logging level
//
// [HH:MM:SS] [WARN] src: msg
//
// IMPORTANT: if you want to write logs to file, open your file in appending mode
func Warn(src string, msg string) {
	if !ensure.NotNil(global) {
		return
	}
	if 2 < global.Level {
		return
	}
	global.chLog <- fill(2, src, msg)
}

// Error is logging msg from src to io.Writer with ERROR logging level
//
// [HH:MM:SS] [ERROR] src: msg
//
// IMPORTANT: if you want to write logs to file, open your file in appending mode
func Error(src string, msg string) {
	if !ensure.NotNil(global) {
		return
	}
	if 3 < global.Level {
		return
	}
	global.chLog <- fill(3, src, msg)
}

// Fatal is logging msg from src to io.Writer with FATAL logging level
//
// [HH:MM:SS] [FATAL] src: msg
//
// IMPORTANT: if you want to write logs to file, open your file in appending mode
func Fatal(src string, msg string) {
	if !ensure.NotNil(global) {
		return
	}
	if 4 < global.Level {
		return
	}
	global.chLog <- fill(4, src, msg)
}

func fill[L int | LoggingLevel](lvl L, src string, msg string) string {
	t := time.Now()
	if reflect.ValueOf(global.Writer).Pointer() == reflect.ValueOf(os.Stdout).Pointer() {
		return fmt.Sprintf("%s[%02d:%02d:%02d] [%s] %s: %s%s",
			color(lvl),
			t.Hour(),
			t.Minute(),
			t.Second(),
			LoggingLevel(lvl),
			src,
			msg,
			ColorReset)
	}
	return fmt.Sprintf("[%02d:%02d:%02d] [%s] %s: %s",
		t.Hour(),
		t.Minute(),
		t.Second(),
		LoggingLevel(lvl),
		src,
		msg)
}

func printer() {
	go func() {
		for str := range global.chLog {
			_, _ = fmt.Fprintln(global.Writer, str)
		}
	}()
}
