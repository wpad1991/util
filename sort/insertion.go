package sort

func InsertionSortInt(arr []int) {

	length := len(arr)

	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				Swap(&arr[j-1], &arr[j])
			} else {
				break
			}
		}
	}
}
