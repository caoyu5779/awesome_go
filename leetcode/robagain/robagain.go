package robagain

func Robagain(nums []int) int {
	size := len(nums)

	switch size {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	return max(robbing(nums[1:]), robbing(nums[:size-1]))
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func robbing(nums []int) int {
	var a, b int

	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(a, b+v)
		}
	}

	return max(a, b)
}
