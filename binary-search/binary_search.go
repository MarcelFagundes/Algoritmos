package binarysearch

// binarysearch finds the index of the target value
// returns -1 if the target is not found
func BinarySearch(values []int, target int) int {
	left, right := 0, len(values)
	for left != right {
		mid := (left + right) / 2
		if values[mid] == target {
			return mid
		} else if values[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
