package sorts

func ShellSort(s []int) {
	for k := len(s) / 2; k > 0; k /= 2 {
		for i := k; i < len(s); i++ {
			l := s[i]
			j := i - k
			for j >= 0 && l < s[j] {
				s[j+k] = s[j]
				j -= k
			}
			s[j+k] = l
		}
	}
}
