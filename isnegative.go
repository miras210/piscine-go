package main

import piscine "./test"
import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
		z01.PrintRune(10)
	} else {
		z01.PrintRune('F')
		z01.PrintRune(10)
	}
}
func main() {

}
