package sorts

func shift(s []int, l int, h int) {
	i := l
	j := 2*i + 1
	k := s[i]
	for j <= h {
		if j < h && s[j] < s[j+1] {
			j++
		}
		if k < s[j] {
			s[i] = s[j]
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
	s[i] = k
}

func HeapSort(s []int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		shift(s, i, len(s)-1)
	}
	for j := len(s) - 1; j > 0; j-- {
		s[0], s[j] = s[j], s[0]
		shift(s, 0, j-1)
	}
}
