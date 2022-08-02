package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	for l.Head != nil {
		l.Head = l.Head.Next

		if l.Head.Next == nil {
			return l.Head.Data
		}
	}
	return l.Head.Data
}
