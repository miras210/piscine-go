package concat

func ConcatParams(args []string) string {
	var ans string
	count := 0
	for range args {
		count++
	}
	for i, str := range args {
		ans += str
		if i != count-1 {
			ans += "\n"
		}
	}
	return ans
}
