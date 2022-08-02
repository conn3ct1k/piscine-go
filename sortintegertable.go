package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := 1; j < len(table); j++ {
			if table[j] < table[j-1] {
				table[j], table[j-1] = table[j-1], table[j]
			}
		}
	}
}
