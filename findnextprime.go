package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}

	for num := nb; ; num-- {
		if isPrime(num) {
			return num
		}
	}
}

func isPrime(nb int) bool {
	if nb < 2 {
		return false
	}

	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
