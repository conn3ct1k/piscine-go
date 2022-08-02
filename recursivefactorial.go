package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb > 1 && nb <= 1000 {
		result = nb * RecursiveFactorial(nb-1)
		if result <= 0 {
			return 0
		}
		return result
	} else if nb == 1 || nb == 0 {
		return 1
	} else {
		return 0
	}
}
