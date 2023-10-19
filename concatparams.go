package piscine

func ConcatParams(args []string) string {
	var ss string
	for i := 0; i < len(args); i++ {
		ss += args[i]
		if i != len(args)-1 {
			ss += "\n"
		}
	}
	return ss
}
