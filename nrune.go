package piscine

func NRune(s string, n int) rune {
	if n > 0 && n <= len(s) {
		n -= 1
		return []rune(s)[n]
	}
	return '\x00'
}
