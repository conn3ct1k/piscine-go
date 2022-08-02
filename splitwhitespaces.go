package piscine

func SplitWhiteSpaces(s string) []string {
	word := ""
	var result []string
	for _, w := range s {
		if w == ' ' && word != "" {
			result = append(result, word)
			word = ""
		} else if w == ' ' {
			continue
		} else {
			word += string(w)
		}
	}
	if word != "" {
		result = append(result, word)
	}

	return result
}
