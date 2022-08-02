package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	asc := false
	desc := false
	for i := 0; i < len(a)-1; i++ {
		if asc {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		} else if desc {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
		if !asc && !desc {
			if f(a[i], a[i+1]) > 0 {
				asc = true
			} else if f(a[i], a[i+1]) < 0 {
				desc = true
			}
		}
	}
	return true
}
