package main

import (
	"os"

	"github.com/01-edu/z01"
)

func convert_s_to_in(s string) int {
	runes := []rune(s)
	result := 0
	for i := 0; i < len(runes); i++ {
		result *= 10
		result += int(runes[i] - 48)
	}
	return result
}

func isNumeric(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if !(runes[i] >= 48 && runes[i] <= 57) {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) == 1 {
		return
	}
	args := os.Args[1:]
	code := 96
	if args[0] == "--upper" {
		if len(os.Args) == 2 {
			return
		}
		code -= 32
		args = args[1:]
	}
	for i := 0; i < len(args); i++ {
		if isNumeric(args[i]) {
			if convert_s_to_in(args[i]) > 0 && convert_s_to_in(args[i]) <= 25 {
				z01.PrintRune(rune(convert_s_to_in(args[i]) + code))
			} else {
				z01.PrintRune(' ')
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
