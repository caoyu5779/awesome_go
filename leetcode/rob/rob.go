package rob

func Rob(nums []int) int {
	n := len(nums)

	if n < 2 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}
	//input []int {1,2,3,1}
	var a , b int
	for  i,v := range nums {
		if i % 2 == 0 {
			a = max(a+v, b)
			//a = 2
			//a = 2+1 =3
		} else {
			b = max(a,b+v)
			//b = 1
			//b = 1+3 = 4
		}
	}

	return max(a,b) // 4
}

func max(i,j int) int {
	if i > j {
		return i
	}

	return j
}


