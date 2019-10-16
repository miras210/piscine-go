package main

import "github.com/01-edu/z01"

func main() {

	var a rune = 122
	for i := a; i >= 97; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}