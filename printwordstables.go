package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, w := range a {
		for _, c := range w {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')

	}
}
