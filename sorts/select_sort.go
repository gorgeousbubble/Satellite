package sorts

func SelectSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		k := i
		for j := i + 1; j < len(s); j++ {
			if s[k] > s[j] {
				k = j
			}
		}
		if k != i {
			s[i], s[k] = s[k], s[i]
		}
	}
}
