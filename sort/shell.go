package sort

func ShellSortInt(arr []int) {

	length := len(arr)

	for gap := length / 2; gap > 0; gap = gap / 2 {
		if gap%2 == 0 {
			gap++
		}

		for i := 0; i < gap; i++ {
			insertionSortGroup(arr, i, length-1, gap)
		}
	}

}

func insertionSortGroup(arr []int, first int, last int, gap int) {

	for i := first + gap; i <= last; i = i + gap {
		key := arr[i]
		j := 0

		for j = i - gap; j >= first && arr[j] > key; j = j - gap {
			arr[j+gap] = arr[j]
		}

		arr[j+gap] = key
	}

}
