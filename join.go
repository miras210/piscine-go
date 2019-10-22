package join

func Join(strs []string, sep string) string {
	var ans string
	var count int = 0
	for i := range strs {
		count++
	}
	for i := range strs {
		ans += strs[i]
		if i != count - 1 {
			ans += sep
		}
	}
	return ans
}
