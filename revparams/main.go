package main

import "os"
import "github.com/01-edu/z01"

func main() {
	arg := os.Args
	count := 0
	for i := range arg {
		count++
	}
	for i := count - 1; i > 0; i-- {
		for _, char := range arg[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
