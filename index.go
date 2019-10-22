package index

func Index(s string, toFind string) int {
	var ind int = -1
	var count int = 0
	for range toFind {
		count++
	}
	c := 0
	for i, char := range s {
		if char == rune(toFind[c]) {
			ind = i + 1
			c++
		} else {
			ind = -1
			c = 0
		}
		if c == count {
			break
		}
	}
	if ind == -1 {
		return ind
	} else {
		return ind - count
	}
}
