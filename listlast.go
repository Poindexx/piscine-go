package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		cc := l.Tail.Data
		return cc
	}
}
