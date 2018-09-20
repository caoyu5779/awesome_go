package sumOfLeftLeaves

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return SumOfLeftLeaves(root.Right)
	}

	if root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + SumOfLeftLeaves(root.Right)
	}

	return SumOfLeftLeaves(root.Left) + SumOfLeftLeaves(root.Right)
}
