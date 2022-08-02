package piscine

func BasicAtoi2(s string) int {
	runes := []rune(s)
	result := 0
	for i := range runes {
		if runes[i] == '0' || runes[i] == '1' || runes[i] == '2' || runes[i] == '3' || runes[i] == '4' || runes[i] == '5' || runes[i] == '6' || runes[i] == '7' || runes[i] == '8' || runes[i] == '9' {
			result *= 10
			result = result + int(runes[i]-'0')
		} else {
			return 0
		}
	}
	return result
}
