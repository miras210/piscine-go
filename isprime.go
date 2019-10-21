package isprime

func IsPrime(nb int) bool {
	for i := 1; i*i <= nb; i++ {
		if nb%i == 0 {
			return true
		}
	}
	return false
}
