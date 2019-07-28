package sorts

func InsertSort(s []int) {
	for i := 1; i < len(s); i++ {
		k := s[i]
		j := i - 1
		for j >= 0 && k < s[j] {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = k
	}
}
