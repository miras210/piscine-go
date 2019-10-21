package recursivefactorial

func RecursiveFactorial(nb int) int {
	if nb >= 0 && nb <= 50 {
		if nb == 0 {
			return 1
		} else {
			return RecursiveFactorial(nb-1) * nb
		}
	} else {
		return 0
	}
}
