package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	counter := 0
	if l == nil {
		return nil
	}
	for l.Next != nil {
		if counter == pos {
			break
		}
		l = l.Next
		counter++
	}
	if counter != pos {
		return nil
	}
	return l
}
