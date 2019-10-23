package main

import "os"
import "github.com/01-edu/z01"

func strNum(s string) int {
	ans := 0
	for _, char := range s {
		ans *= 10
		ans += int(char) - 48
	}
	return ans
}

func isNum(s string) bool {
	for _, char := range s {
		if !(char >= '0' &&  char <='9') {
			return false
		}
	}
	return true
}

func main() {
	arg := os.Args[1:]
	str := "abcdefghijklmnopqrstuvwxyz"
	flag := "--upper"
	mode := false
	if arg[0] == flag {
		mode = true
		arg = arg[1:]
	}
	for _, s := range arg {
		if isNum(s) && strNum(s) >= 1 && strNum(s) <= 26 {		
			if mode {
				z01.PrintRune(rune(str[strNum(s) - 1] - 32))
			} else {
				z01.PrintRune(rune(str[strNum(s) - 1]))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
}
