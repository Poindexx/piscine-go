package piscine

func StrRev(s string) string {
	runes := []rune(s)
	reversed := make([]rune, len(runes))
	for i, j := len(runes)-1, 0; i >= 0; i, j = i-1, j+1 {
		reversed[j] = runes[i]
	}
	return string(reversed)
}
