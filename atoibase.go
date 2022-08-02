package piscine

func AtoiBase(s string, base string) int {
	for i := 0; i < len(base); i++ {
		if count_c(rune(base[i]), base) > 1 || len(base) == 1 {
			return 0
		}
		if rune(base[i]) == '-' || rune(base[i]) == '+' {
			return 0
		}
	}
	for i := 0; i < len(s); i++ {
		if count_c(rune(s[i]), base) == 0 {
			return 0
		}
	}
	ss := len(base)
	rs := []rune(s)
	result := 0
	for i := 0; i < len(rs); i++ {
		result *= ss
		result += Index(base, string(rs[i]))
	}
	return result
}
