package sorts

func merge(s []int, l int, m int, r int) {
	i := l
	j := m + 1
	var t []int
	for i <= m && j <= r {
		if s[i] <= s[j] {
			t = append(t, s[i])
			i ++
		} else {
			t = append(t, s[j])
			j ++
		}
	}
	if i <= m {
		t = append(t, s[i:m+1]...)
	} else {
		t = append(t, s[j:r+1]...)
	}
	for k := 0; k < len(t); k++ {
		s[l+k] = t[k]
	}
}

func MergeSort(s []int, l int, r int) {
	if l < r {
		m := (l + r) / 2
		MergeSort(s, l, m)
		MergeSort(s, m+1, r)
		merge(s, l, m, r)
	}
}
