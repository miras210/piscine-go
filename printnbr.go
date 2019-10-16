package printnbr

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var a rune = 0
	var num int = n
	var count int = 0
	var ten int = 1
	if n == 0 {
		z01.PrintRune('0')
	} else if n < 0 {
		z01.PrintRune('-')
		n = n - 2 * n
		num = n
		for n > 10 {
			for num > 10 {
				num = num / 10
				count++
			}
			for i := a; i <= 127; i++ {
				if i == rune(num + 48) {
					z01.PrintRune(i)
				}	 
			}
			for j := 0; j < count; j++ {
				ten = ten * 10;
			}
			count = 0
			n = n - ten * num
			ten = 1
			num = n
		}
		z01.PrintRune(rune(n+48))
	} else if n > 0 {
		for n > 10 {
			for num > 10 {
				num = num / 10
				count++
			}
			for i := a; i <= 127; i++ {
				if i == rune(num + 48) {
					z01.PrintRune(i)
				}	 
			}
			for j := 0; j < count; j++ {
				ten = ten * 10;
			}
			count = 0
			n = n - ten * num
			ten = 1
			num = n
		}
		z01.PrintRune(rune(n + 48))
	}
}
