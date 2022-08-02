package piscine

func Abort(a, b, c, d, e int) int {
	n := []int{a, b, c, d, e}
	SortIntegerTable(n)
	return n[2]
}
