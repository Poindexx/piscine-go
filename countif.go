package piscine

func CountIf(f func(string) bool, tab []string) int {
	var booll bool
	bb := 0
	for _, v := range tab {
		booll = f(v)
		if booll {
			bb += 1
		}
	}
	return bb
}
