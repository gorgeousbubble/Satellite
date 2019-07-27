package sorts

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
	fmt.Println("Bubble sort:")
	fmt.Println("Before sort slice:", s)
	BubbleSort(s)
	fmt.Println("After sort slice:", s)
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
		BubbleSort(s)
	}
}
