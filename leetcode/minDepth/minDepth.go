package minDepth

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func MinDepth(root *TreeNode) int {
	switch {
	case root == nil:
		return 0

	case root.Left == nil:
		return 1 + MinDepth(root.Right)
	case root.Right == nil:
		return 1 + MinDepth(root.Left)
	default:
		return 1 + min(MinDepth(root.Left), MinDepth(root.Right))
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
