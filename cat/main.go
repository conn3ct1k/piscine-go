package main

import (
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			str := err.Error()
			os.Stderr.WriteString("ERROR: " + str + "\n")
			os.Exit(1)
		}
		out := bytes_to_str(bytes)
		os.Stderr.WriteString(out)
		return
	}
	if len(args) > 0 {
		for _, f := range args {
			print_file(f)
		}
		return
	}
}

func bytes_to_str(bytes []byte) string {
	result := ""
	for _, b := range bytes {
		result += string(b)
	}
	return result
}

func print_file(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		str := err.Error()
		os.Stderr.WriteString("ERROR: " + str + "\n")
		os.Exit(1)
	}
	out := string(file)
	os.Stderr.WriteString(out)
}
