package searches

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	s := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	fmt.Println("Linear Search:")
	index := LinearSearch(s, 10)
	if index == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Println("Index: ", index)
		fmt.Println("s[", index, "] = ", s[index])
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
		LinearSearch(s, 10)
	}
}
