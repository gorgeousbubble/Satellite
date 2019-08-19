package searches

func InsertSearch(s []int, k int) int {
	l := 0
	h := len(s) - 1
	t := 0
	for l < h {
		t += 1
		m := l + int((h-l)*(k-s[l])/(s[h]-s[l]))
		if k < s[m] {
			h = m - 1
		} else if k > s[m] {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}
