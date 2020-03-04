package logs

import (
	"fmt"
	"os"
	. "satellite/logs"
	"time"
)

var log *MultipleLogWriter
var level = LogLevelDebug

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelEvent
	LogLevelWarning
	LogLevelError
	LogLevelCritical
)

func init() {
	// create a multiple file writer instance
	nw, err := NewMultipleFileWriter("./log/erika.log", 1*1024*1024, 5)
	if err != nil {
		fmt.Println("Error init multiple file logs writer:", err)
		os.Exit(1)
	}
	// copy new writer to global variable
	log = nw
}

func Println(level string, a ...interface{}) {
	log.Writef("%s -%s- ", time.Now().Format("2006-01-02 15:04:05"), level)
	log.Writeln(a...)
}

func Debug(a ...interface{}) {
	if level > LogLevelDebug {
		return
	}
	Println("DEBUG", a...)
}

func Info(a ...interface{}) {
	if level > LogLevelInfo {
		return
	}
	Println("INFO", a...)
}

func Event(a ...interface{}) {
	if level > LogLevelEvent {
		return
	}
	Println("EVENT", a...)
}

func Warning(a ...interface{}) {
	if level > LogLevelWarning {
		return
	}
	Println("WARNING", a...)
}

func Error(a ...interface{}) {
	if level > LogLevelError {
		return
	}
	Println("ERROR", a...)
}

func Critical(a ...interface{}) {
	if level > LogLevelCritical {
		return
	}
	Println("CRITICAL", a...)
}
