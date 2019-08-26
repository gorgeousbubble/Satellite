package logs

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const (
	LevelCritical = iota
	LevelError
	LevelWarning
	LevelTrace
	LevelInfo
	LevelDebug
)

const (
	Ltime = 1 << iota
	Lfile
	Llevel
)

const maxBufPoolSize = 16

var levelPrefix = [LevelDebug + 1]string{"[C]", "[E]", "[W]", "[T]", "[I]", "[D]"}

type Atom int32

type Logger struct {
	level   Atom
	flag    int
	lock    sync.Mutex
	handler Handler
	buflock sync.Mutex
	bufs    [][]byte
	closed  Atom
}

func (i *Atom) Set(n int) {
	atomic.StoreInt32((*int32)(i), int32(n))
}

func (i *Atom) Get() int {
	return int(atomic.LoadInt32((*int32)(i)))
}

func NewLogger(handler Handler, flag int) *Logger {
	r := new(Logger)
	r.level.Set(LevelDebug)
	r.flag = flag
	r.closed.Set(0)
	r.bufs = make([][]byte, 0, 16)
	return r
}

func NewDefault(handler Handler) *Logger {
	return NewLogger(handler, Ltime|Lfile|Llevel)
}

func newStdHandler() *StreamHandler {
	h, _ := NewStreamHandler(os.Stdout)
	return h
}

func (l *Logger) popBuf() []byte {
	l.buflock.Lock()
	var buf []byte
	if len(l.bufs) == 0 {
		buf = make([]byte, 0, 1024)
	} else {
		buf = l.bufs[len(l.bufs)-1]
		l.bufs = l.bufs[0 : len(l.bufs)-1]
	}
	l.buflock.Unlock()
	return buf
}

func (l *Logger) putBuf(buf []byte) {
	l.buflock.Lock()
	if len(l.bufs) < maxBufPoolSize {
		buf = buf[0:0]
		l.bufs = append(l.bufs, buf)
	}
	l.buflock.Unlock()
}

func (l *Logger) Close() {
	if l.closed.Get() == 1 {
		return
	}
	l.closed.Set(1)
	l.handler.Close()
}

func (l *Logger) SetLevel(level int) {
	l.level.Set(level)
}

func (l *Logger) SetHandler(h Handler) {
	if l.closed.Get() == 1 {
		return
	}
	l.lock.Lock()
	if l.handler != nil {
		l.handler.Close()
	}
	l.handler = h
	l.lock.Unlock()
}

func (l *Logger) Output(callDepth int, level int, format string, v ...interface{}) {
	if l.closed.Get() == 1 {
		return
	}
	if l.level.Get() > level {
		return
	}
	var s string
	if format == "" {
		s = fmt.Sprint(v...)
	} else {
		s = fmt.Sprintf(format, v...)
	}
	buf := l.popBuf()
	if l.flag&Ltime > 0 {
		now := time.Now().Format("2006/01/02 15:04:05")
		buf = append(buf, '[')
		buf = append(buf, now...)
		buf = append(buf, "] "...)
	}
	if l.flag&Lfile > 0 {
		_, file, line, ok := runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		} else {
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					file = file[i+1:]
					break
				}
			}
		}
		buf = append(buf, file...)
		buf = append(buf, ':')
		buf = strconv.AppendInt(buf, int64(line), 10)
		buf = append(buf, ' ')
	}
	if l.flag&Llevel > 0 {
		buf = append(buf, '[')
		buf = append(buf, levelPrefix[level]...)
		buf = append(buf, "] "...)
	}
	buf = append(buf, s...)
	if s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}
	l.lock.Lock()
	l.handler.Write(buf)
	l.lock.Unlock()
	l.putBuf(buf)
}

func (l *Logger) Cirtical(v ...interface{}) {
	l.Output(2, LevelCritical, "", v)
}

func (l *Logger) Error(v ...interface{}) {
	l.Output(2, LevelError, "", v)
}

func (l *Logger) Warning(v ...interface{}) {
	l.Output(2, LevelWarning, "", v)
}

func (l *Logger) Trace(v ...interface{}) {
	l.Output(2, LevelTrace, "", v)
}

func (l *Logger) Info(v ...interface{}) {
	l.Output(2, LevelInfo, "", v)
}

func (l *Logger) Debug(v ...interface{}) {
	l.Output(2, LevelDebug, "", v)
}

func (l *Logger) Cirticalf(format string, v ...interface{}) {
	l.Output(2, LevelCritical, format, v)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(2, LevelError, format, v)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.Output(2, LevelWarning, format, v)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Output(2, LevelTrace, format, v)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(2, LevelInfo, format, v)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(2, LevelDebug, format, v)
}
