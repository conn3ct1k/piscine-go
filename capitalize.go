package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	prev_not_ap := true
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 97 && runes[i] <= 122 || runes[i] >= 65 && runes[i] <= 90 || runes[i] >= 48 && runes[i] <= 57 {
			if prev_not_ap && runes[i] >= 97 && runes[i] <= 122 {
				runes[i] = runes[i] - 32
			} else if runes[i] >= 65 && runes[i] <= 90 && !prev_not_ap {
				runes[i] = runes[i] + 32
			}
			prev_not_ap = false
		} else {
			prev_not_ap = true
		}
	}
	return string(runes)
}
