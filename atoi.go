package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	runes := []rune(s)
	result := 0
	minus := false
	plus := false
	if runes[0] == '-' {
		minus = true
	}

	if runes[0] == '+' {
		plus = true
	}

	for i := range runes {
		if minus && i == 0 || plus && i == 0 {
			continue
		}
		if runes[i] == '0' || runes[i] == '1' || runes[i] == '2' || runes[i] == '3' || runes[i] == '4' || runes[i] == '5' || runes[i] == '6' || runes[i] == '7' || runes[i] == '8' || runes[i] == '9' {
			result *= 10
			result = result + int(runes[i]-'0')
		} else {
			return 0
		}
	}
	if minus {
		result = -result
	}
	return result
}
