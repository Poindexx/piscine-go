package piscine

func AlphaCount(s string) int {
	co := 0
	for _, e := range s {
		if (e >= 'A' && e <= 'Z') || (e >= 'a' && e <= 'z') {
			co++
		}
	}
	return co
}
