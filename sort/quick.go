package sort

func QuickSortInt(arr []int) {

	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, left int, right int) {

	if left < right {
		q := partition(arr, left, right)
		quicksort(arr, left, q-1)
		quicksort(arr, q+1, right)
	}
}

func partition(arr []int, left int, right int) int {

	low := left + 1
	high := right
	pivot := arr[left]

	for {
		if low > high {
			break
		}

		for {
			if low > right {
				break
			}
			if arr[low] > pivot {
				break
			}

			low++
		}

		for {
			if left > high {
				break
			}
			if arr[high] <= pivot {
				break
			}

			high--
		}
		if low < high {
			Swap(&arr[low], &arr[high])
		}
	}
	Swap(&arr[left], &arr[high])

	return high
}
