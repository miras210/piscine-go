package join

func Join(strs []string, sep string) string {
	var ans string
	for i := range strs {
		ans += strs[i] + sep
	}
	return ans
}
