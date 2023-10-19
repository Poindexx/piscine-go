package piscine

func ListSize(l *List) int {
	co := 0
	cc := l.Head
	for cc != nil {
		co++
		cc = cc.Next
	}
	return co
}
