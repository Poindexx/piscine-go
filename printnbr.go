package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	} else if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits [20]int
	i := 0

	for n > 0 {
		digits[i] = n % 10
		n /= 10
		i++
	}

	for j := i - 1; j >= 0; j-- {
		z01.PrintRune(rune(digits[j]) + '0')
	}
}
