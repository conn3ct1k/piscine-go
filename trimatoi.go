package piscine

func TrimAtoi(s string) int {
	runes := []rune(s)
	minus := false
	first_digit := false
	var digits []int
	for i := 0; i < len(runes); i++ {
		if runes[i] == '-' && !first_digit {
			minus = true
		}
		if runes[i] >= 48 && runes[i] <= 57 {
			first_digit = true
			digits = append(digits, int(runes[i]-48))
		}
	}
	if len(digits) == 0 {
		return 0
	}
	result := 0
	for i := 0; i < len(digits); i++ {
		result *= 10
		result += digits[i]
	}

	if minus {
		result = -result
	}
	return result
}
