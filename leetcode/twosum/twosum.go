package twosum

func TwoSum(nums []int, target int) []int {
	var result []int

	numsLen := len(nums)

	for i := 0; i < numsLen; i++ {
		for j := i + 1; j < numsLen; j++ {
			if target-nums[j] == nums[i] {
				result = append(result, nums[i])
				result = append(result, nums[j])
			}
		}
	}

	return result
}
