package piscine

func IsAlpha(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if !(runes[i] >= 97 && runes[i] <= 122 || runes[i] >= 65 && runes[i] <= 90 || runes[i] >= 48 && runes[i] <= 57) {
			return false
		}
	}
	return true
}
