package permute

import "fmt"

func Permute(nums []int) [][]int {
	n := len(nums)
	// vector 是一组可能的解答
	vector := make([]int, n)
	// taken[i] == true 表示nums[i] 已经出现在了vector中
	taken := make([]bool, n)

	var ans [][]int

	makePermutation(0,n,nums,vector,taken,&ans)
	return ans
}

func makePermutation(cur, n int,nums, vector []int, taken []bool, ans *[][]int)  {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vector)

		*ans = append(*ans, tmp)
		fmt.Println(*ans)
		return
	}

	for i:= 0; i < n; i ++ {
		if !taken[i] {
			taken[i] = true

			vector[cur] = nums[i]
			makePermutation(cur+1, n, nums, vector, taken, ans)

			taken[i] = false
		}
	}
}


