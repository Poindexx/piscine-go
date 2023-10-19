package piscine

func Unmatch(a []int) int {
	m := make(map[int]int)
	for _, e := range a {
		m[e]++
	}
	for _, r := range a {
		if m[r] == 1 {
			return r
		}
	}
	for _, r := range a {
		if m[r] != 2 {
			return r
		}
	}
	return -1
}
