package main

import (
	"os"

	"github.com/01-edu/z01"
)

func sortParams(args []string) {
	for i := 0; i < len(args); i++ {
		for j := 1; j < len(args); j++ {
			if args[j] < args[j-1] {
				args[j], args[j-1] = args[j-1], args[j]
			}
		}
	}
}

func main() {
	args := os.Args[1:]
	sortParams(args)
	for i := 0; i < len(args); i++ {
		arg := []rune(args[i])
		for j := 0; j < len(arg); j++ {
			z01.PrintRune(arg[j])
		}
		z01.PrintRune('\n')
	}
}
