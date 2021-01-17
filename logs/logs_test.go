package logs

import (
	"os"
	"testing"
)

func TestConsoleLog(t *testing.T) {
	w, err := NewConsoleWriter(os.Stdout)
	if err != nil {
		t.Fatal("Error new console writer:", err)
	}
	l := NewDefaultLogger(w)
	l.Info("hello,world!")
}

func TestFileLog(t *testing.T) {
	w, err := NewFileWriter("./log/satellite.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND)
	if err != nil {
		t.Fatal("Error new file writer:", err)
	}
	l := NewDefaultLogger(w)
	l.Debug("hello,world!")
}

func TestFileLogOutput(t *testing.T) {
	w, err := NewFileWriter("./log/satellite.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND)
	if err != nil {
		t.Fatal("Error new file writer:", err)
	}
	l := NewDefaultLogger(w)
	l.Debug("hello!")
	l.Info("Can you speak English?")
	l.Trace("Welcome!")
	l.Warning("What a nice day!")
	l.Error("Something error has happened!")
	l.Cirtical("Critical Fatal...")
}

func TestMultipleFileLog(t *testing.T) {
	w, err := NewMultipleFileWriter("./log/satellite.log", 1*1024*1024, 5)
	if err != nil {
		t.Fatal("Error new multiple writer:", err)
	}
	l := NewDefaultLogger(w)
	l.Debug("hello,world!")
}

func TestMultipleFileLogOutput(t *testing.T) {
	w, err := NewMultipleFileWriter("./log/satellite.log", 1*1024*1024, 5)
	if err != nil {
		t.Fatal("Error new multiple writer:", err)
	}
	l := NewDefaultLogger(w)
	l.Debug("hello!")
	l.Info("Can you speak English?")
	l.Trace("Welcome!")
	l.Warning("What a nice day!")
	l.Error("Something error has happened!")
	l.Cirtical("Critical Fatal...")
}
