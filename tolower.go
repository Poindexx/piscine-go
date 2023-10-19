package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i, e := range runes {
		if e >= 'A' && e <= 'Z' {
			runes[i] = runes[i] + ('a' - 'A')
		}
	}
	return string(runes)
}
