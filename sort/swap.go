package sort

func Swap(a *int, b *int) {

	c := *a
	*a = *b
	*b = c
	// *a = *a + *b
	// *b = *a - *b
	// *a = *a - *b
}
