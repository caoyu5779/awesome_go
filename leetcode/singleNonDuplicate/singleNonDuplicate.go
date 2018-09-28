package singleNonDuplicate

func SingleNonDuplicate(nums []int) int {
	n := len(nums) - 1
	lo, hi := 0, n-1
	for lo < hi {
		mid := lo + (hi-lo)/2

		if mid%2 == 1 {
			mid--
		}

		if nums[mid] != nums[mid+1] {
			hi = mid
		} else {
			lo = mid + 2
		}
	}

	return nums[lo]
}
