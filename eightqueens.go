package piscine

import "github.com/01-edu/z01"

func Check(qn int, rp int, comb [8]int) bool {
	for i := 0; i < qn; i++ {
		t := comb[i]
		if t == rp || t == rp-(qn-i) || t == rp+(qn-i) {
			return false
		}
	}
	return true
}

func Generate(nb int, comb [8]int) {
	if nb == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(comb[i] + 49))
		}
		z01.PrintRune(rune(10))
	} else {
		for i := 0; i < 8; i++ {
			if Check(nb, i, comb) {
				comb[nb] = i
				Generate(nb+1, comb)
			}
		}
	}
}

func EightQueens() {
	comb := [8]int{0, 0, 0, 0, 0, 0, 0}
	Generate(0, comb)
}
