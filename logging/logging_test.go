package logging

import (
	"log"
	"testing"
)

// TestLoggingPlain function
// test logging package with simple text
func TestLoggingPlain(t *testing.T) {
	log.Println("hello,world!")
}

// TestLoggingWithParameters function
// test logging package with parameters
func TestLoggingWithParameters(t *testing.T) {
	a, b := 3, 5
	log.Println("a =", a)
	log.Println("b =", b)
	log.Printf("a + b = %v\n", a + b)
}

// TestLoggingWithCycle function
// test logging package with cycle
func TestLoggingWithCycle(t *testing.T) {
	for i := 0; i < 10000; i++ {
		log.Println("Count:", i)
	}
}

// BenchmarkLoggingPlain function
func BenchmarkLoggingPlain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Println("hello,world!")
	}
}

// BenchmarkLoggingWithParameters function
func BenchmarkLoggingWithParameters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, b := 3, 5
		log.Println("a =", a)
		log.Println("b =", b)
		log.Printf("a + b = %v\n", a + b)
	}
}

// BenchmarkLoggingWithCycle function
func BenchmarkLoggingWithCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10000; i++ {
			log.Println("hello,world!")
		}
	}
}
