package base

func Pow(base, pow int) int {
	ans := base
	if pow == 0 {
		return 1
	} else if pow == 1 {
		return base
	} else {
		for i := 2; i <= pow; i++ {
			ans *= base
		}
		return ans
	}
}

func FindIndex(F rune, base string) int {
	for i, char := range base {
		if char == F {
			return i
		}
	}
	return -1
}

func Rev(s string) string {
	count := 0
	var ans string
	for range s {
		count++
	}
	for i := count - 1; i >= 0; i-- {
		ans += string(s[i])
	}
	return ans
}

func FindChar(Index int, baseTo string) rune {
	for i, char := range baseTo {
		if i == Index {
			return char
		}
	}
	return 0
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	From := 0
	To := 0
	mul := 0
	ten := 0
	for range baseFrom {
		From++
	}
	for range baseTo {
		To++
	}
	nbr = Rev(nbr)
	for i, char := range nbr {
		mul = FindIndex(char, baseTo)
		ten += mul * Pow(From, i)
	}
	var ans string
	for ten != 0 {
		ans += string(FindChar(ten%To, baseTo))
		ten /= To
	}
	ans = Rev(ans)
	return ans
}
