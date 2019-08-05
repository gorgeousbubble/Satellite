package searches

func LinearSearch(s []int, k int) int {
	for i, v := range s {
		if v == k {
			return i
		}
	}
	return -1
}
