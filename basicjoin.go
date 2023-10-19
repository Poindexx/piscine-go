package piscine

func BasicJoin(elems []string) string {
	var s string
	for i := range elems {
		s += elems[i]
	}
	return s
}
