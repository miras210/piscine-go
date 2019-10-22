package toup

func ToUpper(s string) string {
	var s1 string
	for _,char := range s {
		if char >= 'a' && char <= 'z' {
			s1 += string(char - 32)
		} else {
			s1 += string(char)
		}
	}
	return s1
}
