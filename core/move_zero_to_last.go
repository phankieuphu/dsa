package core

func MoveZeroToLast(arr []int) []int {
	n := len(arr)
	holdIndex := 0

	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[holdIndex], arr[i] = arr[i], arr[holdIndex]
			holdIndex++
		}
	}
	return arr
}
