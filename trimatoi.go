package atoi

func TrimAtoi(s string) int {
	var sign bool = true
	var int num
	for _, char := range s {
		if char >= '0' && char <= '9' {
			num = num * 10
			num = num + int(char) - 48
		}
		if char == '-' {
			sign = false
		}
	}
	if sign {
		return num
	} else {
		return (-1) * num
	}
}
