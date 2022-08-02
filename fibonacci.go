package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 1
	} else if index > 2 {
		return Fibonacci(index-1) + Fibonacci(index-2)
	} else {
		return 1
	}
}
