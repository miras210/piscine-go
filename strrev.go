package strrev

func StrRev(s string) string {
	var count int = -1
	for range s {
		count++
	}
	sA := []byte(s)
	for i := 0; i <= count; i++ {
		sA[i] = s[count-i]
	}
	return string(sA)
}
