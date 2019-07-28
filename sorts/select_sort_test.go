package sorts

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
	fmt.Println("Select sort:")
	fmt.Println("Before sort slice:", s)
	SelectSort(s)
	fmt.Println("After sort slice:", s)
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
		SelectSort(s)
	}
}
