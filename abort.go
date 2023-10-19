package piscine

func Abort(a, b, c, d, e int) int {
	aa := make([]int, 5)
	aa[0] = a
	aa[1] = b
	aa[2] = c
	aa[3] = d
	aa[4] = e
	for i := 0; i <= 4; i++ {
		for j := 0; j < 4-i; j++ {
			if aa[j] > aa[j+1] {
				aa[j], aa[j+1] = aa[j+1], aa[j]
			}
		}
	}
	return aa[2]
}
