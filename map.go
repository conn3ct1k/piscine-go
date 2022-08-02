package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, i := range a {
		result = append(result, f(i))
	}
	return result
}
