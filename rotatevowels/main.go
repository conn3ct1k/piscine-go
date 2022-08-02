package main

import (
	"os"

	"github.com/01-edu/z01"
)

// a, e, i, o, u
func main() {
	args := os.Args[1:]
	str := ""
	var vowels_index []int
	vowels_in_args := false
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}
	for i := 0; i < len(args); i++ {
		t := args[i]
		for j := 0; j < len(t); j++ {
			str += string(t[j])
		}
		str += " "
	}

	for i := 0; i < len(str); i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'u' || str[i] == 'o' || str[i] == 'A' || str[i] == 'E' || str[i] == 'I' || str[i] == 'U' || str[i] == 'O' {
			vowels_in_args = true
			vowels_index = append(vowels_index, i)
		}
	}

	if !vowels_in_args {
		for i := 0; i < len(str); i++ {
			z01.PrintRune(rune(str[i]))
		}
		z01.PrintRune('\n')
		return
	}

	if vowels_in_args {
		runes := []rune(str)
		for i := 0; i < len(vowels_index)/2; i++ {
			runes[vowels_index[i]], runes[vowels_index[len(vowels_index)-1-i]] = runes[vowels_index[len(vowels_index)-1-i]], runes[vowels_index[i]]
		}
		str = string(runes)
		for _, c := range str {
			z01.PrintRune(c)
		}

	}
	z01.PrintRune('\n')
}
