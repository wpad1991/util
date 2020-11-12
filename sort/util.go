package sort

func Swap(a *int, b *int) {

	c := *a
	*a = *b
	*b = c
	// *a = *a + *b
	// *b = *a - *b
	// *a = *a - *b
}

func CheckLinearArray(sortedarray []int) []int {

	empty := make([]int, 0)

	for i := 0; i < len(sortedarray)-1; i++ {
		diff := sortedarray[i+1] - sortedarray[i]
		if diff == 0 {

		} else if diff != 1 {
			for j := 1; j < diff; j++ {
				empty = append(empty, sortedarray[i]+j)
			}
		}
	}

	return empty
}
