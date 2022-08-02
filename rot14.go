package piscine

func Rot14(s string) string {
	result := ""
	alpha_lower := "abcdefghijklmnopqrstuvwxyz"
	alpha_upper := ToUpper(alpha_lower)
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			result += string(alpha_lower[(c-97+14)%26])
		} else if c >= 'A' && c <= 'Z' {
			result += string(alpha_upper[(c-65+14)%26])
		} else {
			result += string(c)
		}
	}
	return result
}
