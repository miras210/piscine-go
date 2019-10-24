package split

func SplitWhiteSpaces(str string) []string {
	count := 0
	for _, char := range str {
		if char == '\n' || char == ' ' || char == '\t' {
			count++
		}
	}
	arr := make([]string, count+1)
	i := 0
	next := false
	for i != count {
		for _, char := range str {
			if !(char == '\n' || char == ' ' || char == '\t') {
				arr[i] += string(char)
				next = true
			} else {
				if next {
					i++
					next = false
				}
			}
		}
	}
	return arr
}
