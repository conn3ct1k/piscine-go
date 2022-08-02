package main

import (
	"os"
)

var isErr int

func main() {
	args := os.Args[3:]
	nStr := os.Args[2]
	n := 0

	for i := 0; i < len(nStr); i++ {
		n *= 10
		n += int(rune(nStr[i]) - 48)
	}

	for i, f := range args {
		print_file(f, n, i)
	}

	if isErr == 1 {
		os.Exit(1)
	}
}

func print_file(filename string, n int, idx int) {
	file, err := os.ReadFile(filename)
	out := ""
	if err != nil {
		str := err.Error()
		os.Stderr.WriteString(str + "\n")
		isErr = 1
	}
	if err == nil {
		if idx > 0 {
			os.Stderr.WriteString("\n")
		}
		os.Stderr.WriteString("==> " + filename + " <==")
		os.Stderr.WriteString("\n")
		if len(file)-1 < n {
			out = string(file)
		} else {
			out = string(file[len(file)-n:])
		}
		os.Stderr.WriteString(out)
	}
}
