package main

import "github.com/01-edu/z01"

func PrintComb() {
	var a int = 48
	var b int = 49
	var c int = 50
	for i:=0; i <= 6; i++ {
		for j:=i; j <=7; j++ {
			for k:=j; k <=7; k++ {
				z01.PrintRune(rune(a+i))
				z01.PrintRune(rune(b+j))
				z01.PrintRune(rune(c+k))
				z01.PrintRune(44)
				z01.PrintRune(32)
				
			}
		}
	}
	z01.PrintRune('7')
	z01.PrintRune('8')
	z01.PrintRune('9')
}

func main() {
	PrintComb()
}