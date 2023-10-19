package piscine

func Join(strs []string, sep string) string {
	var s string
	for i := range strs[:len(strs)-1] {
		s += strs[i]
		s += string(sep)
	}
	s += string(strs[len(strs)-1])
	return s
}
