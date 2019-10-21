package iterativefactorial

func IterativeFactorial(nb int) int {
	if nb >= 0 {
		ans := 1
		for i := 2; i <= nb; i++ {
			ans = ans * i
		}
		return ans
	} else {
		return 0
	}
}
