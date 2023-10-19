package piscine

func Rot14(s string) string {
	var a string
	for _, e := range s {
		if e >= 'A' && e < 'M' || e >= 'a' && e < 'm' {
			a += string(e + 14)
		} else if e >= 'M' && e <= 'Z' || e >= 'm' && e <= 'z' {
			a += string(e - 12)
		} else {
			a += string(e)
		}
	}
	return string(a)
}
