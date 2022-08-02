package piscine

func BasicJoin(elems []string) string {
	result := ""
	for i := range elems {
		result += elems[i]
	}
	return result
}
