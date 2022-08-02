package piscine

func Split(s, sep string) []string {
	isSep := false
	word := ""
	var result []string
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] {
			for j := range sep {
				if len(s) < i+j {
					isSep = false
					break
				}
				if s[i+j] != sep[j] {
					isSep = false
					break
				}
				isSep = true
			}
		}
		if isSep {
			result = append(result, word)
			word = ""
			i += len(sep) - 1
			isSep = false
			continue
		}
		word += string(s[i])
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
