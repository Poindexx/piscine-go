package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var prev *NodeL
	for iter := l.Head; iter != nil; iter = iter.Next {
		if iter.Data == data_ref {
			if iter == l.Head {
				l.Head = iter.Next
			} else {
				prev.Next = iter.Next
			}
		}
		prev = iter
	}
	for iter := l.Head; iter != nil; iter = iter.Next {
		if iter.Data == data_ref {
			if iter == l.Head {
				l.Head = iter.Next
			} else {
				prev.Next = iter.Next
			}
		}
		prev = iter
	}
}
