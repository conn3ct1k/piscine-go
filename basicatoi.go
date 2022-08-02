package piscine

func BasicAtoi(s string) int {
	runes := []rune(s)
	result := 0
	for i := range runes {
		result *= 10
		result = result + int(runes[i]-'0')
	}
	return result
}
