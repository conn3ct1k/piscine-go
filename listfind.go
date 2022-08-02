package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	cn := l.Head
	for cn != nil {
		if comp(cn.Data, ref) {
			return &cn.Data
		}
		cn = cn.Next
	}
	return nil
}
