package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}
	if l == nil {
		return newNode
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = newNode
	sorted := true
	for sorted {
		sorted = false
		current := l
		for current != nil && current.Next != nil {
			if current.Data > current.Next.Data {
				current.Data, current.Next.Data = current.Next.Data, current.Data
				sorted = true
			}
			current = current.Next
		}
	}
	return l
}
