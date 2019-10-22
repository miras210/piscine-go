package base

import "github.com/01-edu/z01"

func print(base string, num int) {
	for i, char := range base {
		if i == num {
			z01.PrintRune(char)
		}
	}
}

func checkBase(base string) bool {
	count := 0
	var str string
	for _, char := range base {
		if char == '-' || char == '+' {
			return false
		}
		count++
		for _, char1 := range str {
			if char1 == char {
				return false
			}
		}
		str += string(char)
	}
	return true
}

func two(nbr int, base string, sign bool) {
	var ans string
	count := 0
	if sign == false {
		nbr = (-1)*nbr
	}
	for nbr != 0 {
		if nbr % 2 == 1 {
			ans += string(base[1])
		} else {
			ans += string(base[0])
		}
		nbr /= 2
	}
	for range ans {
		count ++
	}
	for i := count - 1; i >= 0; i-- {
		z01.PrintRune(rune(ans[i]))
	}
}

func PrintNbrBase(nbr int, base string) {
	count := 0
	power := 0
	var sign bool = true
	if checkBase(base) {
		for range base {
			count++
		}
		if nbr < 0 {
			sign = false
			z01.PrintRune('-')
		}
		if count == 2 {
			two(nbr, base, sign)
		} else {
			for nbr != 0 {
				num := 0
				if sign == false {
					nbr = (-1)*nbr
					num = nbr
					sign = true
				} else {
					num = nbr
				}
				power = 0
				for num >= count {
					num /= count
					power++
				}
				print(base, num)
				minus := 1
				for i := 0; i < power; i++ {
					minus *= count
				}
				nbr -= minus*num
			}
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}
