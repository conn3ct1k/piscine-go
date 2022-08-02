package piscine

func AppendRange(min, max int) []int {
	var result []int
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}
