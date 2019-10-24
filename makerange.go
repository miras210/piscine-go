package make

func MakeRange(min, max int) []int {
	size := max - min
	arr := make([]int, size)
	i := 0
	for min < max {
		arr[i] = min
		min++
		i++
	}
	return arr
}
