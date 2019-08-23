package logs

import (
	"sync"
	"sync/atomic"
)

const (
	LevelCritical = iota
	LevelError
	LevelWarning
	LevelTrace
	LevelInfo
	LevelDebug
)

var levelPrefix = [LevelDebug + 1]string{"[C]", "[E]", "[W]", "[T]", "[I]", "[D]"}

type Atom int32

type Handler interface {
	Write(p []byte) (n int, err error)
	Close() error
}

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

//func NewDefault(handler Handler) *Logger {
//	return NewLogger(handler, )
//}
