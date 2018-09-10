package sortColors

func SortColors(nums []int) []int {
	length := len(nums)

	if length == 0 {
		return nil
	}

	temp := nums[0]
	nums[0] = 1

	i,j,k := 0,1, length -1
	
	for j <= k {
		switch  {
		case nums[j] < 1:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case 1 < nums[j]:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		default:
			j++
		}
	}

	switch temp {
	case 0:
		nums[i] = temp
	case 2:
		nums[k] = temp
	}
	return nums
}
