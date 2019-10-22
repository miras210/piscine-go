package index

func Index(s string, toFind string) int {
	var ind int = -1
	var count int = 0
	for index, char := range s {
		if char == rune(toFind[count]) {
			ind = index + 1
			count++
		} else {
			ind = -1
			count = 0
		}
		if count >= len(toFind) {
			break
		}
	}
	if ind == -1 {
		return ind
	} else {
		return ind - count
	}
}
