package printcomb2

import "github.com/01-edu/z01"

func PrintComb2() {
	var a1 rune = '0'
	var a2 rune = '0'
	var b1 rune = '0'
	var b2 rune = '1'
	for d := a1; d <= 57; d++ {
		for v := a2; v < 57; v++ {
			for i := b1; i <= 57; i++ {
				for j := b2; j <= 57; j++ {
					if d == '9' && v == '8' && i == '9' && j == '9' {
						break
					} else if d < i && j != '0' {
						z01.PrintRune(d)
						z01.PrintRune(v)
						z01.PrintRune(32)
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(44)
						z01.PrintRune(32)
					} else if d == i {
						if v < j && j != '0' {
							z01.PrintRune(d)
							z01.PrintRune(v)
							z01.PrintRune(32)
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(44)
							z01.PrintRune(32)
						}
					}
				}
			}
		}
	}
	z01.PrintRune('9')
	z01.PrintRune('8')
	z01.PrintRune(32)
	z01.PrintRune('9')
	z01.PrintRune('9')
	z01.PrintRune(10)
}
