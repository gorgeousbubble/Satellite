package sorts

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
	fmt.Println("Merge sort:")
	fmt.Println("Before sort slice:", s)
	MergeSort(s, 0, 9)
	fmt.Println("After sort slice:", s)
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
		MergeSort(s, 0, 9)
	}
}
