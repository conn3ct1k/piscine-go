package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	abase := AtoiBase(nbr, baseFrom)

	result := ReturnNbrBase(abase, baseTo)
	return StrRev(result)
}

func ReturnNbrBase(nbr int, base string) string {
	var arrInt []rune
	arrBase := []rune(base)

	if nbr < 0 {
		nbr *= -1
	}

	for nbr > 0 {
		arrInt = append(arrInt, arrBase[nbr%len(arrBase)])
		nbr /= len(arrBase)
	}
	return string(arrInt)
}
