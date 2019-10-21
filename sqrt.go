package sqrt

func Sqrt(nb int) int {
	ans := 0
	for i := 1; i <= nb; i ++ {
		if i*i == nb {
			ans = i
			return ans
		}
	}
	return ans
}
