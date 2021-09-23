package logger

import (
	"fmt"
	"log"
)

type Severity int

const (
	DEB Severity = iota
	INF
	WRN
	ERR
)

const (
	yellow = "\033[90;43m"
	red    = "\033[97;41m"
	blue   = "\033[97;44m"
	reset  = "\033[0m"
)

func SetOutput()  {

}

func (s Severity) String() string {
	switch s {
	case DEB:
		return "DEB"
	case INF:
		return "INF"
	case WRN:
		return "WRN"
	case ERR:
		return "ERR"
	default:
		return "unknown severity"
	}
}

func Info(msg string, vars ...interface{}) {
	printLog(INF, msg, blue, vars)
}

func Warn(msg string, vars ...interface{}) {
	printLog(WRN, msg, yellow, vars)
}

func Debug(msg string, vars ...interface{}) {
	printLog(DEB, msg, reset, vars)
}

func Error(msg string, vars ...interface{}) {
	printLog(ERR, msg, red, vars)
}

func Panic(err error) {
	log.Panic(err)
}

func printLog(severity Severity, msg string, color string, vars ...interface{}) {
	log.Printf(fmt.Sprintf("%s %s %s %s", color, severity, reset, msg), vars...)
}
