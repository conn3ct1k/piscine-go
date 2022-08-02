package piscine

func IsLower(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if !(runes[i] >= 97 && runes[i] <= 122) {
			return false
		}
	}
	return true
}
