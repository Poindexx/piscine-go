package piscine

func UltimateDivMod(a *int, b *int) {
	q := *a / *b
	w := *a % *b
	*a = q
	*b = w
}
