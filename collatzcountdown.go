package piscine

func CollatzCountdown(start int) int {
	co := 0
	if start <= 0 {
		return -1
	}
	for {
		if start%2 == 0 {
			start = start / 2
			co += 1
		} else {
			start = start*3 + 1
			co += 1
		}
		if start == 1 {
			break
		}
	}
	return co
}
