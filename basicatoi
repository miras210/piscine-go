package basicatoi

func BasicAtoi(s string) int {
	var ans int = 0
	var ansR rune = 0
	var check bool = false
	var count int = -1
	for _, char := range s {
		if check {
			for i := ansR; i <= 9; i++ {
				count++
				if i == rune(char)-48 {
					break
				}
			}
			ans = ans * 10
			ans = ans + count
			count = -1
		} else {
			if char != '0' {
				check = true
				for i := ansR; i <= 9; i++ {
					count++
					if i == rune(char)-48 {
						break
					}
				}
				ans = count
				count = -1
			}
		}
	}
	return ans
}
