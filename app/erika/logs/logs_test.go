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

func BenchmarkPrintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Println("DEBUG", "hello,world!")
	}
}

func BenchmarkDebug(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Debug("hello,world!")
	}
}

func BenchmarkInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info("This is a message~")
	}
}

func BenchmarkEvent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Event("One Http Request###")
	}
}

func BenchmarkWarning(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Warning("Don't touch it!")
	}
}

func BenchmarkError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Error("Error call function!!")
	}
}

func BenchmarkCritical(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Critical("Critical Error!!!")
	}
}
