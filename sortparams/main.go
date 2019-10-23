package main

import "os"
import "github.com/01-edu/z01"

func main() {
	arg := os.Args
	count := 0
	for i, char := range arg {
		if i != 0 {
			for j := ' '; j <= '~'; j++ {
				if j == char {
					z01.PrintRune(char)
				}
			}
			z01.PrintRune('\n')
		}
	}
}
