package piscine

func StrRev(s string) string {
	var result string = ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}
