package logs

import (
	"fmt"
	"github.com/spf13/pflag"
	"time"
)

var (

)

const (
	InfoLevel = "INFO"
	WarnLevel = "WARN"
	ErrorLevel = "ERROR"
)

func InitLogs(flags *pflag.FlagSet) {

}

func Infof(format string, a ...interface{}) {
	format = fmt.Sprintf("[%5s][%s] %s", InfoLevel, time.Now().Format("2006-01-02 15:04:05"), format)
	fmt.Printf(format, a...)
}

func Warnf(format string, a ...interface{}) {
	format = fmt.Sprintf("[%5s][%s] %s", WarnLevel, time.Now().Format("2006-01-02 15:04:05"), format)
	fmt.Printf(format, a...)
}

func Errorf(format string, a ...interface{}) {
	format = fmt.Sprintf("[%5s][%s] %s", ErrorLevel, time.Now().Format("2006-01-02 15:04:05"), format)
	fmt.Printf(format, a...)
}