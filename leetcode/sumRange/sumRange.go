package sumRange

import "fmt"

type NumArray struct {
	dp []int
}


func Constructor(nums []int) NumArray {
	size := len(nums)

	dp := make([]int, size+1)

	for i:=1; i <= size; i++ {
		dp[i] = dp[i-1]+nums[i-1]
		fmt.Println(dp)
	}

	return NumArray{dp:dp}
}


func (na *NumArray) SumRange(i int, j int) int {
	return na.dp[j+1] - na.dp[i]
}

