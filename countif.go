package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, w := range tab {
		if f(w) {
			counter++
		}
	}
	return counter
}
