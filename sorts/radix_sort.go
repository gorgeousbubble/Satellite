package sorts

import (
	"math"
	"strconv"
)

func RadixSort(s []int) {
	max := s[0]
	for i := 0; i < len(s); i++ {
		if max < s[i] {
			max = s[i]
		}
	}
	for j := 0; j < len(strconv.Itoa(max)); j++ {
		bin := make([][]int, 10)
		for k := 0; k < len(s); k++ {
			n := s[k] / int(math.Pow(10, float64(j))) % 10
			bin[n] = append(bin[n], s[k])
		}
		m := 0
		for p := 0; p < len(bin); p++ {
			for q := 0; q < len(bin[p]); q++ {
				s[m] = bin[p][q]
				m++
			}
		}
	}
}
