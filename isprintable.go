package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if !(runes[i] >= 32) {
			return false
		}
	}
	return true
}
