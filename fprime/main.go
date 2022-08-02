package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		n, _ := strconv.Atoi(args[0])
		del := 2
		for n > 1 {
			if n%del == 0 {
				fmt.Print(del)
				n /= del
				if n > 1 {
					fmt.Print("*")
				}
				del--
			}
			del++
		}
		fmt.Println()
	}
}
