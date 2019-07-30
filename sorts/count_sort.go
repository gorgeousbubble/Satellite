package sorts

func CountSort(s []int) {
	max := s[0]
	for i := 1; i < len(s); i++ {
		if max < s[i] {
			max = s[i]
		}
	}
	arr := make([]int, max+1)
	for j := 0; j < len(s); j++ {
		arr[s[j]]++
	}
	k := 0
	for m, n := range arr {
		for p := 0; p < n; p++ {
			s[k] = m
			k++
		}
	}
}
