package iterativepower

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		ans := 0
		for i := 1; i <= power; i++ {
			ans = ans * nb
		}
		return ans
	}
}
