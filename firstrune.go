package fRune

func FirstRune(s string) rune {
	var a rune = '0'
	for _, char := range s {
		a = rune(char)
		break
	}
	return a
}
