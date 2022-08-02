package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
		l.Tail = l.Head
	} else {
		l.Tail.Next = &NodeL{Data: data}
		l.Tail = l.Tail.Next
	}
}
