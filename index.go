package piscine

func Index(s string, toFind string) int {
	index := -1
	for i := range s {
		if s[i] == toFind[0] {
			for k := range toFind {
				if len(s) <= i+k {
					index = -1
					break
				}
				if s[i+k] != toFind[k] {
					index = -1
					break
				}
				index = i
			}
		}
		if index != -1 {
			return index
		}
	}
	return index
}
