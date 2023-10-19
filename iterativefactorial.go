package piscine

func IterativeFactorial(nb int) int {
	if nb > 20 || nb < 0 {
		return 0
	}
	r := 1
	for i := 2; i <= nb; i++ {
		r *= i
	}
	return r
}
