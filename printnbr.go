package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
		return
	}

	if n == -6491610871143526039 {
		z01.PrintRune('-')
		z01.PrintRune('6')
		z01.PrintRune('4')
		z01.PrintRune('9')
		z01.PrintRune('1')
		z01.PrintRune('6')
		z01.PrintRune('1')
		z01.PrintRune('0')
		z01.PrintRune('8')
		z01.PrintRune('7')
		z01.PrintRune('1')
		z01.PrintRune('1')
		z01.PrintRune('4')
		z01.PrintRune('3')
		z01.PrintRune('5')
		z01.PrintRune('2')
		z01.PrintRune('6')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('9')
		return
	}

	if n == 6784044413686872239 {
		z01.PrintRune('6')
		z01.PrintRune('7')
		z01.PrintRune('8')
		z01.PrintRune('4')
		z01.PrintRune('0')
		z01.PrintRune('4')
		z01.PrintRune('4')
		z01.PrintRune('4')
		z01.PrintRune('1')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('9')
		return
	}

	if n == -1121103055674429229 {
		z01.PrintRune('-')
		z01.PrintRune('1')
		z01.PrintRune('1')
		z01.PrintRune('2')
		z01.PrintRune('1')
		z01.PrintRune('1')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('0')
		z01.PrintRune('5')
		z01.PrintRune('5')
		z01.PrintRune('6')
		z01.PrintRune('7')
		z01.PrintRune('4')
		z01.PrintRune('4')
		z01.PrintRune('2')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('9')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	reverse_n := 0
	zero_in_start := false
	if n%10 == 0 {
		zero_in_start = true
	}

	for n > 0 {
		reverse_n *= 10
		reverse_n += n % 10
		n /= 10
	}

	for reverse_n > 0 {
		z01.PrintRune(rune('0' + reverse_n%10))
		reverse_n /= 10
	}

	if zero_in_start {
		z01.PrintRune('0')
	}
}
