package utils

import "testing"

func TestGenRSAKey2File(t *testing.T) {
	err := GenRSAKey2File(1024)
	if err != nil {
		t.Errorf("Error Generate RSA Key to File.")
	}
}

func TestGenRSAKey2Memory(t *testing.T) {
	var private []byte
	var public []byte
	err := GenRSAKey2Memory(&private, &public, 1024)
	if err != nil {
		t.Errorf("Error Generate RSA Key to Memory.")
	}
}

func BenchmarkGenRSAKey2File(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := GenRSAKey2File(1024)
		if err != nil {
			b.Errorf("Error Generate RSA Key to File.")
		}
	}
}

func BenchmarkGenRSAKey2Memory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var private []byte
		var public []byte
		err := GenRSAKey2Memory(&private, &public, 1024)
		if err != nil {
			b.Errorf("Error Generate RSA Key to Memory.")
		}
	}
}
