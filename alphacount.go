package piscine

func AlphaCount(s string) int {
	runes := []rune(s)
	counter := 0
	for c := range runes {
		if runes[c] >= 97 && runes[c] <= 122 || runes[c] >= 65 && runes[c] <= 90 {
			counter++
		}
	}
	return counter
}
