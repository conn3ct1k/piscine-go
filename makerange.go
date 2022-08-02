package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	result := make([]int, max-min)
	if len(result) > 0 {
		for i := 0; i < len(result); i++ {
			result[i] = i + min
		}
	}
	return result
}
