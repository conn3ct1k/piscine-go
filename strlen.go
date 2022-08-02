package piscine

func StrLen(s string) int {
	runes := []rune(s)
	counter := 0
	for i := range runes {
		counter = i
	}
	counter++
	return counter
}
