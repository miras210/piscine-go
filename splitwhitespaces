package split

func SplitWhiteSpaces(str string) []string {
	count := 0
	single := true
	for _, char := range str {
		if !(char == '\n' || char == ' ' || char == '\t') {
			single = true
		} else {
			if single {
				count++
				single = false
			}
		}
	}
	arr := make([]string, count+1)
	i := 0
	duble := true
	for i != count {
		for _, char := range str {
			if !(char == '\n' || char == ' ' || char == '\t') {
				arr[i] += string(char)
				duble = true
			} else {
				if duble {
					i++
					duble = false
				}
			}
		}
	}
	return arr
}
