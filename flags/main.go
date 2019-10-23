package main

import "os"
import "github.com/01-edu/z01"

func print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func Insert(add string, s string) string {
	var Nadd string
	var Ns string
	beg := false
	for _, char := range add {
		if beg {
			if char >= '!' && char <= '~' {
				Nadd += string(char)
			} else {
				Nadd += string(' ')
			}
		}
		if char == '=' {
			beg = true
		}
	}
	for _, char := range s {
		if char >= '!' && char <= '~' {
			Ns += string(char)
		} else {
			Ns += string(' ')
		}
	}
	return Ns + Nadd
}

func Order(s string) string {
	var ans string
	var i rune
	for i := 0; i <= 127; i++ {
		for _, char := range s {
			if i == char {
				ans += string(char)
			}
		}
	}
	return ans
}

func Help(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	arg := os.Args[1:]
	var ans string
	isFlag := false
	isEmpty := true
	text := "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument."
	count := 0
	for range arg {
		count++
	}
	if count == 0 {
		Help(text)
	} else if count == 1 {
		if arg[0] == "-h" || arg[0] == "--help" {
			Help(text)
		} else {
			print(arg[0])
		}
	} else {
		for i := range arg {
			for _, char := range arg[i] {
				if isFlag {
					if char == 'i' {
						ans = Insert(arg[i], arg[count-1])
						isEmpty = false
						isFlag = false
						break
					} else if char == 'o' {
						if isEmpty == false {
							ans = Order(ans)
						} else {
							ans = Order(arg[count-1])
						}
						isFlag = false
						break
					} else if char == 'h' {
						Help(text)
						isFlag = false
						break
					}
				}
				if char == '-' {
					isFlag = true
				}
			}
		}
		print(ans)
	}
	z01.PrintRune('\n')
}
