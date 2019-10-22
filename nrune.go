package nRune

func NRune(s string) rune {
	var a rune = '0'
	for _, char := range s {
		a = char
	}
	return a
}
