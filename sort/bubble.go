package sort

func BubbleSortInt(arr []int) {

	length := len(arr)

	for i := 0; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if j+1 < length && arr[j] > arr[j+1] {
				Swap(&arr[j], &arr[j+1])
			}
		}
	}
}
