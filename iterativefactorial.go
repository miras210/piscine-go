package iterativefactorial

func IterativeFactorial(nb int) int {
	if nb == 1 {
		return 1
	} else if nb >= 1 {
		return IterativeFactorial(nb-1) * nb
	} else {
		return 0
	}
}
