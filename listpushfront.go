package piscine

func ListPushFront(l *List, data interface{}) {
	l.Head = &NodeL{Data: data, Next: l.Head}
}
