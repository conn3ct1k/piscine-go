package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, c := range args {
		if c == "galaxy 01" || c == "galaxy" || c == "01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
