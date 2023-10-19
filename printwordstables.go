package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := 0; i <= len(a)-1; i++ {
		argument := a[i]
		for j := 0; j <= len(argument)-1; j++ {
			z01.PrintRune(rune(argument[j]))
		}
		z01.PrintRune('\n')
	}
}
