package piscine

func IsPrintable(s string) bool {
	for _, e := range s {
		if !(e >= 32 && e <= 126) {
			return false
		}
	}
	return true
}
