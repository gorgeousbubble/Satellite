package sorts

func BucketSort(s []int, num int) {
	min, max := s[0], s[0]
	for i := 0; i < len(s); i++ {
		if min > s[i] {
			min = s[i]
		}
		if max < s[i] {
			max = s[i]
		}
	}
	bucket := make([][]int, num)
	for j := 0; j < len(s); j++ {
		n := (s[j] - num) / ((max - min + 1) / num)
		bucket[n] = append(bucket[n], s[j])
		k := len(bucket[n]) - 2
		for k >= 0 && s[j] < bucket[n][k] {
			bucket[n][k+1] = bucket[n][k]
			k--
		}
		bucket[n][k+1] = s[j]
	}
	o := 0
	for p, q := range bucket {
		for t := 0; t < len(q); t++ {
			s[o] = bucket[p][t]
			o++
		}
	}
}
