package findTargetSumWays

import "fmt"

func FindTargetSumWays(nums []int, S int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum < S {
		return 0
	}

	if (sum+S)%2 == 1 {
		return 0
	}

	return calcSumWays(nums, (sum+S)/2)
}

/*
	nums 分为两个部分 N P
	P前面放 "－"
	N前面放 "+"
	可得 sum(P)-sum(N)=target
	sum(P)+sum(N)+sum(P)-sum(N)=target+sum(P)+sum(N)
	2 * sum(P) = target + sum(nums)
	sum(P)=(target+sum(nums))/2
*/
func calcSumWays(nums []int, target int) int {
	//dp [i] 表示 nums中 和为i的子集个数
	dp := make([]int, target+1)

	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
			fmt.Println(dp)
		}
	}

	return dp[target]
}
