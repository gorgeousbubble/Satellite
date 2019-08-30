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
