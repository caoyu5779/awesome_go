package searchRange

func SearchRange(nums []int, target int) []int {
	index := search(nums, target)

	if index == -1 {
		return []int{-1, -1}
	}

	first := index

	//找到头
	for {
		f := search(nums[:first], target)
		if f == -1 {
			break
		}

		first = f
	}

	last := index

	//找到尾
	for {
		l := search(nums[last+1:], target)
		if l == -1 {
			break
		}

		last += l + 1
	}

	return []int{first, last}
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	var middle int
	for low <= high {
		middle = (low + high) / 2

		switch {
		case nums[middle] < target:
			low = middle + 1
		case nums[middle] > target:
			high = middle - 1
		default:
			return middle
		}
	}

	return -1
}
