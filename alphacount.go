package alphacount

import "github.com/01-edu/z01"

func AlphaCount(str string) int {
	count := 0
	for _,char := range str {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
		} 
	}
	return count
}
