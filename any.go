package piscine

func Any(f func(string) bool, a []string) bool {
	for _, w := range a {
		if f(w) {
			return true
		}
	}
	return false
}
