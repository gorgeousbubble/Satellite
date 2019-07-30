package sorts

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
	fmt.Println("Bucket sort:")
	fmt.Println("Before sort slice:", s)
	BucketSort(s, 1)
	fmt.Println("After sort slice:", s)
}

func BenchmarkBucketSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{1, 3, 8, 2, 4, 9, 7, 5, 6, 0}
		BucketSort(s, 1)
	}
}
