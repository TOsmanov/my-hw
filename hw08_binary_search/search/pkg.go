package search

func BinarySearch(arr []int, key int) (int, bool) {
	min := 0
	high := len(arr) - 1

	for min <= high {
		mid := (min + high) / 2
		midValue := arr[mid]

		switch {
		case midValue < key:
			min = mid + 1
		case midValue > key:
			high = mid - 1
		default:
			return mid, true
		}
	}
	return -1, false
}
