package compare

func Compare(a, b string) int {
	count1 := 0
	count2 := 0
	for range a {
		count1++
	}
	for range b {
		count2++
	}
	for i, j := 0, 0; i < count1 && j < count2; i, j = i+1, j+1 {
		if a[i] > b[j] {
			return 1
		} else if a[i] < b[j] {
			return -1
		}
	}
	return 0
}
