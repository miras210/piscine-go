package main

import "os"
import "github.com/01-edu/z01"

func main() {
	arg := os.Args
	for i := range arg {
		if i != 0 {
			for _, char := range arg[i] {
				z01.PrintRune(char)
			}
		}
		z01.PrintRune('\n')
	}
}
