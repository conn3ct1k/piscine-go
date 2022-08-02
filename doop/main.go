package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 2 {
		first := args[0]
		op := args[1][0]
		second := args[2]
		if !(IsNumeric(first)) || !(IsNumeric(second)) {
			return
		}
		a := parse_int(first)
		b := parse_int(second)
		if op == '%' && b == 0 {
			fmt.Println("No modulo by 0")
			return
		} else if op == '/' && b == 0 {
			fmt.Println("No division by 0")
			return
		}
		result := 0
		if a > 9223372036854775750 || b > 9223372036854775750 {
			return
		}
		switch op {
		case '*':
			result = a * b
			if a == 0 || result/a == b {
				fmt.Println(result)
			}
		case '/':
			result = a / b
			fmt.Println(result)
		case '%':
			result = a % b
			fmt.Println(result)
		case '-':
			result = a - b
			if (result < a) && (b > 0) {
				fmt.Println(result)
			}
		case '+':
			result = a + b
			if (result > a) && (b > 0) {
				fmt.Println(result)
			}
		}
	}
}

func IsNumeric(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if i == 0 && runes[0] == '-' {
			continue
		}
		if !(runes[i] >= 48 && runes[i] <= 57) {
			return false
		}
	}
	return true
}

func parse_int(s string) int {
	minus := false
	result := 0
	for i, c := range s {
		if s[0] == '-' && i == 0 {
			minus = true
			continue
		}
		result *= 10
		result += int(c - '0')
	}
	if minus {
		result = -result
	}
	return result
}
