package piscine

func IsAlpha(s string) bool {
	for _, e := range s {
		if !((e >= 'A' && e <= 'Z') || (e >= 'a' && e <= 'z') || (e >= '0' && e <= '9')) {
			return false
		}
	}
	return true
}
