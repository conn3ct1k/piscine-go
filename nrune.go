package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if len(s) < n || n < 1 {
		return 0
	}
	for i := range runes {
		if i == n-1 {
			return rune(runes[i])
		}
	}
	return 0
}
