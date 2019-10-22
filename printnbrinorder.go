package order

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	s := Itoa(n)
	var s1 string
	for i := '0'; i <= '9'; i++ {
		for _, char := range s {
			if i == char {
				z01.PrintRune(char)
			}
		}
	}
}
