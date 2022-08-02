package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}

	setPoint(&points)

	for _, c := range "x = " {
		z01.PrintRune(c)
	}
	printnbr(points.x)
	for _, c := range ", y = " {
		z01.PrintRune(c)
	}
	printnbr(points.y)
	z01.PrintRune('\n')
}

func printnbr(a int) {
	n := 0
	for a > 0 {
		n *= 10
		n += a % 10
		a /= 10
	}

	for n > 0 {
		z01.PrintRune(rune('0' + n%10))
		n /= 10
	}
}
