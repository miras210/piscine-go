package basicatoi2

func BasicAtoi2(s string) int {
	var ans int = 0
	for _, char := range s {
		if int(char) - 48 >= 0 && int(char) - 48 <= 9 {
			ans = 10 * ans
			ans = ans + int(char) - 48
		} else {
			return 0
		}
	}
	return ans
}
