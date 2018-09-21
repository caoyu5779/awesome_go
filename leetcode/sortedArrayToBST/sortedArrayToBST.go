package sortedArrayToBST

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0{
		return nil
	}

	mid := len(nums)/2
	return &TreeNode{
		Val : nums[mid],
		Left: SortedArrayToBST(nums[:mid]),
		Right: SortedArrayToBST(nums[mid+1:]),
	}

}
