package isBalanced

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func IsBalanced(root *TreeNode) bool {
	_,isBalanced := recur(root)
	return isBalanced
}

func recur(root *TreeNode) (int,bool) {
	if root == nil {
		return 0, true
	}

	leftDepth , leftIsBalanced := recur(root.Left)
	rightDepth, rightIsBalanced := recur(root.Right)

	if leftIsBalanced && rightIsBalanced && -1 <= leftDepth - rightDepth && leftDepth - rightDepth <= 1 {
		return max(leftDepth, rightDepth) + 1,true
	}

	return 0, false
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
