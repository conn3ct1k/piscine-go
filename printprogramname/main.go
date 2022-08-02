package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := []rune(os.Args[0])
	for i := 2; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	z01.PrintRune('\n')
}
