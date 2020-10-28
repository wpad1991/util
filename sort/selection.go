package sort

func SelectionSortInt(arr []int) {

	length := len(arr)

	for i := 0; i < length; i++ {

		min := i

		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j

			}
		}

		if min != i {
			Swap(&arr[min], &arr[i])
		}
	}
}
