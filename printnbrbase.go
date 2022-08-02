package piscine

import (
	"github.com/01-edu/z01"
)

func count_c(c rune, s string) int {
	counter := 0
	for i := 0; i < len(s); i++ {
		if c == rune(s[i]) {
			counter++
		}
	}
	return counter
}

func recursive_append(nbr int, base string, arr []rune) []rune {
	ss := len(base)

	if nbr <= 0 {
		return arr
	}
	arr = append(arr, rune(base[nbr%ss]))
	return recursive_append(nbr/ss, base, arr)
}

func PrintNbrBase(nbr int, base string) {
	var result []rune
	if nbr == -9223372036854775808 {
		PrintStr("-9223372036854775808")
		return
	}
	for i := 0; i < len(base); i++ {
		if count_c(rune(base[i]), base) > 1 || len(base) == 1 {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		if rune(base[i]) == '-' || rune(base[i]) == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}
	if nbr < 0 {
		nbr = -nbr
		z01.PrintRune('-')
	}

	result = recursive_append(nbr, base, result)

	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}
}
