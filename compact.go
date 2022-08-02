package piscine

func Compact(ptr *[]string) int {
	result := []string{}
	for _, el := range *ptr {
		if el != "" {
			result = append(result, el)
		}
	}
	*ptr = result
	return len(result)
}
