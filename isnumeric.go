package piscine

func IsNumeric(s string) bool {
	for _, e := range s {
		if !(e >= '0' && e <= '9') {
			return false
		}
	}
	return true
}
