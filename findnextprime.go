package findnextprime

func FindNextPrime(nb int) int {
	if nb > 1 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return FindNextPrime(nb+1)
			}
		}
		return nb
	} else {
		return nb
	}
}
