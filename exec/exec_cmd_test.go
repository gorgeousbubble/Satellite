package exec

import (
	"fmt"
	"testing"
)

// TestExecCmd function
func TestExecCmd(t *testing.T) {
	s := "cmd"
	r, err := ExecCmd(s)
	if err != nil {
		t.Fatal("Error exec command:", err)
	}
	fmt.Println(r)
}

// TestExecCmd2 function
func TestExecCmd2(t *testing.T) {
	s := "mstsc"
	r, err := ExecCmd(s)
	if err != nil {
		t.Fatal("Error exec command:", err)
	}
	fmt.Println(r)
}

// TestExecCmd3 function
func TestExecCmd3(t *testing.T) {
	s := "dxdiag"
	r, err := ExecCmd(s)
	if err != nil {
		t.Fatal("Error exec command:", err)
	}
	fmt.Println(r)
}

// BenchmarkExecCmd function
func BenchmarkExecCmd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "cmd"
		_, err := ExecCmd(s)
		if err != nil {
			b.Fatal("Error exec command:", err)
		}
	}
}
