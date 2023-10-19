package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
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
