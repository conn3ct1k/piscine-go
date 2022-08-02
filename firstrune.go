package piscine

func FirstRune(s string) rune {
	for _, c := range s {
		return rune(c)
	}
	return 0
}
