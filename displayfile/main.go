package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	file, err := os.Open(args[0])
	data := make([]byte, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file.Read(data)
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			break
		}
		fmt.Print(string(data[i]))
	}
}
