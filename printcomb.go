package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := i; j <= '9'; j++ {
			for l := j; l <= '9'; l++ {
				if i < j && j < l {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(l)
					if i == '7' && j == '8' && l == '9' {
						break
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
