package make

func MakeRange(min, max int) []int {
	size := max - min
	if size <= 0 {
		var arr []int
		return arr
	} else {
		arr := make([]int, size)
		i := 0
		for min < max {
			arr[i] = min
			min++
			i++
		}
		return arr
	}
}
