package hasPathSum

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func HasPathSum(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	target -= root.Val

	if root.Left == nil && root.Right == nil {
		return target == 0
	}

	return HasPathSum(root.Left, target) || HasPathSum(root.Right, target)
}
