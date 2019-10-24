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
	for i != count {
		for _, char := range str {
			if !(char == '\n' || char == ' ' || char == '\t') {
				arr[i] += string(char)
			} else {
				i++
			}
		}
	}
	return arr
}
