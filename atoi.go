package atoi

func Atoi(s string) int {
	var ans int = 0
	var sign bool = false;
	if s == "" {
		return 0
	} else {
		if s[0] == '-' || s[0] == '+' {
			sign = true;
		}
		if s[0] == '-' {
			for _, char := range s {
				if sign {
					sign = false
					continue
				}
				if int(char) - 48 >= 0 && int(char) - 48 <= 9 {
					ans = 10 * ans
					ans = ans + int(char) - 48
				} else {
					return 0
				}
			}
			return (-1)*ans
		} else if s[0] == '+' {
			for _, char := range s {
				if sign {
					sign = false
					continue
				}
				if int(char) - 48 >= 0 && int(char) - 48 <= 9 {
					ans = 10 * ans
					ans = ans + int(char) - 48
				} else {
					return 0
				}
			}
			return ans
		} else {
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
	}
}
