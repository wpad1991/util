package sort

import (
	"testing"
)

/*
Test 1 : 순차적인 숫자
Test 2 : 한개의 숫자
Test 3 : 중복되는 수가 포함되어 있는 숫자
Test 4 : -값을 포함하는 숫자
*/

func TestCheckLinearArray(t *testing.T) {
	arr := []int{1, 2, 4, 6, 9, 100}

	arr2 := CheckLinearArray(arr)

	for _, val := range arr2 {
		println(val)
	}
}

func TestQuickSortInt(t *testing.T) {
	qwe1 := []int{2, 5, 23, 1, 52}
	qwe2 := []int{2, 5, 23, 1, 52, 0, 12, 34, 53, 55, 66, 32, 1, 5}
	qwe3 := []int{2}
	qwe4 := []int{2, 5, 23, 1, 52, 0, 12}
	qwe5 := []int{1, 2, 3, 2, 1}
	qwe6 := []int{1, 2, 3, 2, 1, 0, 0, 2, 3, 3, 4, 4, 4, 2, 2, 1, 1}
	qwe7 := []int{1, -2, 3, -2, 1, 0, 0, 2, -3, 3, 4, -4, 4, 2, 2, 1, 1}
	qwe8 := []int{2, 5, 23, 1, -52, 0, 12, 34, -53, 55, 66, 32, -1, 5}

	QuickSortInt(qwe1)
	QuickSortInt(qwe2)
	QuickSortInt(qwe3)
	QuickSortInt(qwe4)
	QuickSortInt(qwe5)
	QuickSortInt(qwe6)
	QuickSortInt(qwe7)
	QuickSortInt(qwe8)

	println(checksort(qwe1))
	println(checksort(qwe2))
	println(checksort(qwe3))
	println(checksort(qwe4))
	println(checksort(qwe5))
	println(checksort(qwe6))
	println(checksort(qwe7))
	println(checksort(qwe8))
}

func TestShellSortInt(t *testing.T) {
	qwe1 := []int{2, 5, 23, 1, 52}
	qwe2 := []int{2, 5, 23, 1, 52, 0, 12, 34, 53, 55, 66, 32, 1, 5}
	qwe3 := []int{2}
	qwe4 := []int{2, 5, 23, 1, 52, 0, 12}
	qwe5 := []int{1, 2, 3, 2, 1}
	qwe6 := []int{1, 2, 3, 2, 1, 0, 0, 2, 3, 3, 4, 4, 4, 2, 2, 1, 1}
	qwe7 := []int{1, -2, 3, -2, 1, 0, 0, 2, -3, 3, 4, -4, 4, 2, 2, 1, 1}
	qwe8 := []int{2, 5, 23, 1, -52, 0, 12, 34, -53, 55, 66, 32, -1, 5}

	ShellSortInt(qwe1)
	ShellSortInt(qwe2)
	ShellSortInt(qwe3)
	ShellSortInt(qwe4)
	ShellSortInt(qwe5)
	ShellSortInt(qwe6)
	ShellSortInt(qwe7)
	ShellSortInt(qwe8)

	println(checksort(qwe1))
	println(checksort(qwe2))
	println(checksort(qwe3))
	println(checksort(qwe4))
	println(checksort(qwe5))
	println(checksort(qwe6))
	println(checksort(qwe7))
	println(checksort(qwe8))
}

func TestInsertion(t *testing.T) {
	qwe1 := []int{2, 5, 23, 1, 52}
	qwe2 := []int{2, 5, 23, 1, 52, 0, 12, 34, 53, 55, 66, 32, 1, 5}
	qwe3 := []int{2}
	qwe4 := []int{2, 5, 23, 1, 52, 0, 12}
	qwe5 := []int{1, 2, 3, 2, 1}
	qwe6 := []int{1, 2, 3, 2, 1, 0, 0, 2, 3, 3, 4, 4, 4, 2, 2, 1, 1}
	qwe7 := []int{1, -2, 3, -2, 1, 0, 0, 2, -3, 3, 4, -4, 4, 2, 2, 1, 1}
	qwe8 := []int{2, 5, 23, 1, -52, 0, 12, 34, -53, 55, 66, 32, -1, 5}

	InsertionSortInt(qwe1)
	InsertionSortInt(qwe2)
	InsertionSortInt(qwe3)
	InsertionSortInt(qwe4)
	InsertionSortInt(qwe5)
	InsertionSortInt(qwe6)
	InsertionSortInt(qwe7)
	InsertionSortInt(qwe8)

	println(checksort(qwe1))
	println(checksort(qwe2))
	println(checksort(qwe3))
	println(checksort(qwe4))
	println(checksort(qwe5))
	println(checksort(qwe6))
	println(checksort(qwe7))
	println(checksort(qwe8))

}

func TestSelection(t *testing.T) {
	qwe1 := []int{2, 5, 23, 1, 52}
	qwe2 := []int{2, 5, 23, 1, 52, 0, 12, 34, 53, 55, 66, 32, 1, 5}
	qwe3 := []int{2}
	qwe4 := []int{2, 5, 23, 1, 52, 0, 12}
	qwe5 := []int{1, 2, 3, 2, 1}
	qwe6 := []int{1, 2, 3, 2, 1, 0, 0, 2, 3, 3, 4, 4, 4, 2, 2, 1, 1}
	qwe7 := []int{1, -2, 3, -2, 1, 0, 0, 2, -3, 3, 4, -4, 4, 2, 2, 1, 1}
	qwe8 := []int{2, 5, 23, 1, -52, 0, 12, 34, -53, 55, 66, 32, -1, 5}

	SelectionSortInt(qwe1)
	SelectionSortInt(qwe2)
	SelectionSortInt(qwe3)
	SelectionSortInt(qwe4)
	SelectionSortInt(qwe5)
	SelectionSortInt(qwe6)
	SelectionSortInt(qwe7)
	SelectionSortInt(qwe8)

	println(checksort(qwe1))
	println(checksort(qwe2))
	println(checksort(qwe3))
	println(checksort(qwe4))
	println(checksort(qwe5))
	println(checksort(qwe6))
	println(checksort(qwe7))
	println(checksort(qwe8))

}

func TestBubble(t *testing.T) {
	qwe1 := []int{2, 5, 23, 1, 52}
	qwe2 := []int{2, 5, 23, 1, 52, 0, 12, 34, 53, 55, 66, 32, 1, 5}
	qwe3 := []int{2}
	qwe4 := []int{2, 5, 23, 1, 52, 0, 12}
	qwe5 := []int{1, 2, 3, 2, 1}
	qwe6 := []int{1, 2, 3, 2, 1, 0, 0, 2, 3, 3, 4, 4, 4, 2, 2, 1, 1}
	qwe7 := []int{1, -2, 3, -2, 1, 0, 0, 2, -3, 3, 4, -4, 4, 2, 2, 1, 1}
	qwe8 := []int{2, 5, 23, 1, -52, 0, 12, 34, -53, 55, 66, 32, -1, 5}

	BubbleSortInt(qwe1)
	BubbleSortInt(qwe2)
	BubbleSortInt(qwe3)
	BubbleSortInt(qwe4)
	BubbleSortInt(qwe5)
	BubbleSortInt(qwe6)
	BubbleSortInt(qwe7)
	BubbleSortInt(qwe8)

	println(checksort(qwe1))
	println(checksort(qwe2))
	println(checksort(qwe3))
	println(checksort(qwe4))
	println(checksort(qwe5))
	println(checksort(qwe6))
	println(checksort(qwe7))
	println(checksort(qwe8))

}

func arrprint(arr []int) {
	for _, value := range arr {
		println(value)
	}
}

func checksort(arr []int) bool {

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
