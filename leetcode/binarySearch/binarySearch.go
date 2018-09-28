package binarySearch

func BinarySearch(nums []int, key int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + (end - 1)) / 2
		if nums[mid] == key {
			return mid
		} else if nums[mid] > key {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
