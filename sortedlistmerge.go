package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	for list1 := n1; list1 != nil; list1 = list1.Next {
		if list1.Next == nil {
			for list2 := n2; list2 != nil; list2 = list2.Next {
				list1.Next = list2
				list1 = list2
			}
		}
	}
	sorted := true
	for sorted {
		sorted = false
		current := n1
		for current != nil && current.Next != nil {
			if current.Data > current.Next.Data {
				current.Data, current.Next.Data = current.Next.Data, current.Data
				sorted = true
			}
			current = current.Next
		}
	}
	return n1
}
