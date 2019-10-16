package printnbr

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var a rune = 0
	var num int = n
	var save int = n
	var numR rune = 48
	var count int = 0
	var ten int = 1
	if n == 0 {
		z01.PrintRune('0')
	} else if n < 0 {
		if n == -9223372036854775808 {
			z01.PrintRune('-')
			z01.PrintRune('9')
			z01.PrintRune('2')
			z01.PrintRune('2')
			z01.PrintRune('3')
			z01.PrintRune('3')
			z01.PrintRune('7')
			z01.PrintRune('2')
			z01.PrintRune('0')
			z01.PrintRune('3')
			z01.PrintRune('6')
			z01.PrintRune('8')
			z01.PrintRune('5')
			z01.PrintRune('4')
			z01.PrintRune('7')
			z01.PrintRune('7')
			z01.PrintRune('5')
			z01.PrintRune('8')
			z01.PrintRune('0')
			z01.PrintRune('8')
		} else {
			z01.PrintRune('-')
			n = n * (-1)
			num = n
			save = n
			for n >= 10 {
				for num >= 10 {
					num = num / 10
					count++
				}
				save = num
				for num != 0 {
					numR = numR + 1
					num--
				}
				for i := a; i <= 127; i++ {
					if i == numR {
						z01.PrintRune(i)
					}
				}
				for j := 0; j < count; j++ {
					ten = ten * 10
				}
				count = 0
				n = n - ten*save
				if n < ten/10 {
					z01.PrintRune('0')
				}
				save = n
				ten = 1
				num = n
				numR = 48
			}
		}
		numR = 48
		for i := 0; i <= n; i++ {
			numR = numR + 1
		}
		z01.PrintRune(numR - 1)
	} else if n > 0 {
		for n >= 10 {
			for num >= 10 {
				num = num / 10
				count++
			}
			save = num
			for num != 0 {
				numR = numR + 1
				num--
			}
			for i := a; i <= 127; i++ {
				if i == numR {
					z01.PrintRune(i)
				}
			}
			for j := 0; j < count; j++ {
				ten = ten * 10
			}
			count = 0
			n = n - ten*save
			if n < ten/10 {
				z01.PrintRune('0')
			}
			save = n
			ten = 1
			num = n
			numR = 48
		}
		numR = 48
		for i := 0; i <= n; i++ {
			numR = numR + 1
		}
		z01.PrintRune(numR - 1)
	}
}
