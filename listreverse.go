package piscine

func ListReverse(l *List) {
	n := l.Head
	var next *NodeL
	var prev *NodeL

	for n != nil {
		next = n.Next
		n.Next = prev
		prev = n
		n = next

	}
	l.Head = prev
}
