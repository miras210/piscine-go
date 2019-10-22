package cap

func Capitalize(s string) string {
	var s1 string
	var prev rune = 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			if !((prev >= 'a' && prev <= 'z') || (prev >= 'A' && prev <= 'Z') || (prev >= '0' && prev <= '9')) {
				if char >= 'A' && char <= 'Z' {
					s1 += string(char)
				} else {
					s1 += string(char - 32)
				}
			} else {
				if char >= 'A' && char <= 'Z' {
					s1 += string(char + 32)
				} else {
					s1 += string(char)
				}
			}
		} else if char >= 'A' && char <= 'Z' {
			s1 += string(char + 32)
		} else {
			s1 += string(char)
		}
		prev = char
	}
	return s1
}
