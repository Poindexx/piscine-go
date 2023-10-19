package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	co := 0
	cc := l
	for co < pos {
		if cc == nil {
			return nil
		}
		co++
		cc = cc.Next
	}
	return cc
}
