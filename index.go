package piscine

func Index(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {
		rr := true

		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				rr = false
				break
			}
		}
		if rr {
			return i
		}
	}
	return -1
}
