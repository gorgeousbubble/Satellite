package sorts

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
	fmt.Println("Quick sort:")
	fmt.Println("Before sort slice:", s)
	QuickSort(s, 0, 9)
	fmt.Println("After sort slice:", s)
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
		QuickSort(s, 0, 9)
	}
}
