package maxDepth

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
