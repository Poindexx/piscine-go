package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i, e := range runes {
		if e >= 'a' && e <= 'z' {
			runes[i] = runes[i] - ('a' - 'A')
		}
	}
	return string(runes)
}
