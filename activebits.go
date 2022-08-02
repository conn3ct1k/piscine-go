package piscine

func ActiveBits(n int) int {
	counter := 0
	for _, c := range dec_to_bin(n) {
		if c == '1' {
			counter++
		}
	}
	return counter
}

func dec_to_bin(n int) string {
	s := ""
	for n > 0 {
		s += string('0' + n%2)
		n /= 2
	}
	return s
}
