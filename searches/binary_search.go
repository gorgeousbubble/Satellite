package searches

import "math"

func BinarySearch(s []int, k int) int {
	l, r, m := 1, len(s), 0
	for {
		m = int(math.Floor(float64((l + r) / 2)))
		if s[m] > k {
			r = m - 1
		} else if s[m] < k {
			l = m + 1
		} else {
			break
		}
		if l > r {
			m = -1
			break
		}
	}
	return m
}
