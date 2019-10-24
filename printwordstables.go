package table

import "github.com/01-edu/z01"

func print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func PrintWordsTables(table []string) {
	count := 0
	for range table {
		count++
	}
	for i, str := range table {
		print(str)
		if i != count-1 {
			z01.PrintRune('\n')
		}
	}
}
