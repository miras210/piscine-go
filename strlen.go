package strlen

func StrLen(str string) int {
	var count int = 0
	for i, _ := range str {
		count = i
	}
	return count + 1
}
