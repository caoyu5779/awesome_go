package wiggleMaxLength

func WiggleMaxLength(nums []int) int {
	var res int
	size := len(nums)

	checkFunc := func(n int) func(int) {
		init := 1
		if n < 0 {
			init = -1
		}

		res = 2

		return func(x int) {
			var new int

			switch  {
			case x < 0:
				new = -1
			case x > 0:
				new = 1
			default:
				return
			}

			if init * new < 0 {
				res++
				init = new
			}
		}
	}

	var check func(int)

	var i = 1

	for i < size && nums[i] - nums[i-1] == 0{
		i ++
	}

	if i == size {
		return 1
	}

	check = checkFunc(nums[i]-nums[i-1])

	for i < size {
		check(nums[i] - nums[i-1])
		i++
	}

	return res
}
