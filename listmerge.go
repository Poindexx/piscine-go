package piscine

func ListMerge(l1 *List, l2 *List) {
	for list2 := l2.Head; list2 != nil; list2 = list2.Next {
		if l1.Head == nil {
			l1.Head = list2
			l1.Tail = list2
		} else {
			l1.Tail.Next = list2
			l1.Tail = list2
		}
	}
}
