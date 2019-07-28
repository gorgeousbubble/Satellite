package sorts

import (
	"math/rand"
	"time"
)

func QuickSort(s []int, l int, r int) {
	if l >= r {
		return
	}
	i := l
	j := r
	rand.Seed(time.Now().Unix())
	x := rand.Intn(r-l) + l
	s[i], s[x] = s[x], s[i]
	k := s[i]
	for i < j {
		for i < j && s[j] >= k {
			j--
		}
		s[i] = s[j]
		for i < j && s[i] <= k {
			i++
		}
		s[j] = s[i]
	}
	s[i] = k
	QuickSort(s, l, i-1)
	QuickSort(s, i+1, r)
}
