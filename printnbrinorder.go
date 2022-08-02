package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	var result []int
	if n == 0 {
		z01.PrintRune(48)
	}
	for n > 0 {
		result = append(result, n%10)
		n /= 10
	}
	SortIntegerTable(result)
	for i := 0; i < len(result); i++ {
		z01.PrintRune(rune(48 + result[i]))
	}
}
