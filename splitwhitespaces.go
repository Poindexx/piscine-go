package piscine

func SplitWhiteSpaces(s string) []string {
	delimiter := " "
	var substrings []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == delimiter[0] {
			if i+1 < len(s) && s[i:i+len(delimiter)] == delimiter {
				substrings = append(substrings, s[start:i])
				start = i + len(delimiter)
				i += len(delimiter) - 1
			}
		}
	}
	substrings = append(substrings, s[start:])
	for i := 0; i < len(substrings); i++ {
		if len(substrings[i]) == 0 {
			substrings = append(substrings[:i], substrings[i+1:]...)
		}
	}
	return substrings
}
