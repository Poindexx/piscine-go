package piscine

func Any(f func(string) bool, a []string) bool {
	var booll bool
	bb := false
	for _, v := range a {
		booll = f(v)
		if booll {
			bb = true
		}
	}
	return bb
}
