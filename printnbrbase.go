package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		qq := []rune(base)
		for i := 0; i < len(qq)-1; i++ {
			for j := i + 1; j < len(qq); j++ {
				if qq[i] == qq[j] || qq[i] == '-' || qq[j] == '-' || qq[i] == '+' || qq[j] == '+' {
					z01.PrintRune('N')
					z01.PrintRune('V')
					return
				}
			}
		}
		if nbr < 0 {
			z01.PrintRune('-')
		}
		var aa []rune
		for nbr != 0 {
			last := nbr % len(base)
			if last < 0 {
				last *= -1
			}
			aa = append(aa, qq[last])
			nbr /= len(base)
		}
		for i := len(aa) - 1; i >= 0; i-- {
			z01.PrintRune(aa[i])
		}
	}
}
