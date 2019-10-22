package print

func IsPrintable(str string) bool {
	for _, char := range str {
		if char < 32 || char >126 {
			return false
		}
	}
	return true
}
