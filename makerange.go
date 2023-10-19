package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		var arr []int
		return arr
	} else if min == 0 {
		var arr []int
		arr = make([]int, 1)
		return arr
	} else {
		arr := make([]int, max-min)
		for i := min; i < max; i++ {
			arr[i-min] = i
		}
		return arr
	}
}
