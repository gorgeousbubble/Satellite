package sorts

import (
	"fmt"
	"testing"
)

func TestCountSort(t *testing.T) {
	s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
	fmt.Println("Count sort:")
	fmt.Println("Before sort slice:", s)
	CountSort(s)
	fmt.Println("After sort slice:", s)
}

func BenchmarkCountSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
		CountSort(s)
	}
}
