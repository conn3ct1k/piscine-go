package piscine

func ConcatParams(args []string) string {
	result := ""
	for i, w := range args {
		result += w
		if i != len(args)-1 {
			result += "\n"
		}
	}
	return result
}
