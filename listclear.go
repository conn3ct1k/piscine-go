package piscine

func ListClear(l *List) {
	for l.Head != nil {
		l.Head.Data = nil
		l.Head = l.Head.Next
	}
}
