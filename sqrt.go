package piscine

func Sqrt(nb int) int {
	r := 1
	for i := 1; i <= 100000; i++ {
		r = i * i
		if r == nb {
			return i
		}
	}
	return 0
}
