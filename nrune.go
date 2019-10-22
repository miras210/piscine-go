package nRune

func NRune(s string, n int) rune {
	var a rune = '0'
	for i, char := range s {
		if i == n-1 {
			a = char
		}
	}
	return '\x00'
}
