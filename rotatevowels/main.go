package main

import "os"
import "github.com/01-edu/z01"

func print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func IsVowel(c rune, s string) bool {
	for _, char := range s {
		if char == c {
			return true
		}
	}
	return false
}

func main() {
	arg := os.Args[1:]
	vowels := "AaUuEeOoIi"
	var StrVow string
	var ans string
	countV := 0
	count := 0
	for range arg {
		count++
	}
	if count == 0 {
		z01.PrintRune('\n')
	} else {
		for i := range arg {
			for _, char := range arg[i] {
				if IsVowel(char, vowels) {
					countV++
					StrVow += string(char)
				}
			}
		}
		if countV == 0 {
			for i, str := range arg {
				print(str)
				if i != count-1 {
					z01.PrintRune(' ')
				}
			}
		} else {
			for i := range arg {
				for _, char := range arg[i] {
					if IsVowel(char, vowels) {
						ans += string(StrVow[countV-1])
						countV--
					} else {
						ans += string(char)
					}
				}
				if i != count-1 {
					ans += " "
				}
			}
			print(ans)
		}
		z01.PrintRune('\n')
	}
}
