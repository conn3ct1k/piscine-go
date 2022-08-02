package main

import (
	"fmt"
	"os"
)

func help_message() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func main() {
	args := os.Args
	strappend := []rune{}
	inprogress := true
	if len(args) == 1 {
		help_message()
		inprogress = false
	} else if (args[1] == "--help" || args[1] == "-h") && inprogress {
		help_message()
		inprogress = false
	} else {
		first := []rune(args[1])
		ins := []rune("--insert=")
		mainstr := []rune{}
		remain := true
		order := false
		for i := 0; i <= 8; i++ {
			if ins[i] != first[i] {
				remain = false
				break
			}
		}
		index := 9
		if first[0] == '-' && first[1] == 'i' && first[2] == '=' {
			remain = true
			index = 3
		}
		if remain {
			for i := index; i < len(first); i++ {
				strappend = append(strappend, first[i])
			}
			if args[2] == "-o" || args[2] == "--order" {
				order = true
				mainstr = []rune(args[3])
			} else {
				mainstr = []rune(args[2])
			}
		} else {
			if args[1] == "-o" || args[1] == "--order" {
				order = true
				mainstr = []rune(args[2])
			} else {
				mainstr = []rune(args[1])
			}
		}
		mainstr = append(mainstr, strappend...)
		if order {
			for i := 0; i < len(mainstr)-1; i++ {
				for j := 0; j < len(mainstr)-1-i; j++ {
					if mainstr[j] > mainstr[j+1] {
						mainstr[j], mainstr[j+1] = mainstr[j+1], mainstr[j]
					}
				}
			}
		}
		fmt.Println(string(mainstr))
	}
}
