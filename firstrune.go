package fRune

func FirstRune(s string) rune {
	for _, char := range s {
		return rune(char)
	}
}
