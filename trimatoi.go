package piscine

func TrimAtoi(s string) int {
	sign := 1
	n := 0
	digit := false
	for _, char := range s {
		if digit && char == 48 {
			n = n*10 + int(char-'0')
		}
		if char > 48 && char <= 57 {
			digit = true
			n = n*10 + int(char-'0')
		}
		if char == '-' && !digit {
			sign = -1
		}
	}
	return n * sign
}
