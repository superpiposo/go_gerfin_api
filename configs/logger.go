package configs

import (
	"log"
	"os"
)

type Logger struct {
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func NewLogger(prefix string) *Logger {
	return &Logger{
		Debug:   log.New(os.Stdout, prefix+" [DEBUG] ", log.LstdFlags),
		Info:    log.New(os.Stdout, prefix+" [INFO]  ", log.LstdFlags),
		Warning: log.New(os.Stdout, prefix+" [WARN]  ", log.LstdFlags),
		Error:   log.New(os.Stdout, prefix+" [ERROR] ", log.LstdFlags),
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Info.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.Warning.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Error.Printf(format, v...)
}
