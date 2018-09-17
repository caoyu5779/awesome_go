package maxSubArray

import "fmt"

func MaxSubArray(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	temp := nums[0]
	//-2
	max := temp

	for i := 1; i < n; i++{
		if temp < 0 {
			temp = nums[i]
			fmt.Println("-" ,temp)
		} else {
			temp += nums[i]
			fmt.Println("+" ,temp)
		}

		if max < temp {
			max = temp
		}
	}

	return max
}
