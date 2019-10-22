package atoi

func TrimAtoi(s string) int {
	var sign bool = true
	var num int = 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			num = num * 10
			num = num + int(char) - 48
		} else if char == '-' && num == 0 {
			sign = false
		}
	}
	if sign {
		return num
	} else {
		return (-1) * num
	}
}
