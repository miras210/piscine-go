package order

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var s string
	s += rune(n%10) + 48
	for n != 0 {
		n = n / 10
		s += rune(n%10) + 48
	}
	for i := '0'; i <= '9'; i++ {
		for _, char := range s {
			if i == char {
				z01.PrintRune(char)
			}
		}
	}
}
