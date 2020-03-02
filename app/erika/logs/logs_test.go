package logs

import "testing"

func TestPrintln(t *testing.T) {
	Println("DEBUG", "hello,world!")
}

func TestDebug(t *testing.T) {
	Debug("hello,world!")
}

func TestInfo(t *testing.T) {
	Info("This is a message~")
}

func TestEvent(t *testing.T) {
	Event("One Http Request###")
}

func TestWarning(t *testing.T) {
	Warning("Don't touch it!")
}

func TestError(t *testing.T) {
	Error("Error call function!!")
}

func TestCritical(t *testing.T) {
	Critical("Critical Error!!!")
}
