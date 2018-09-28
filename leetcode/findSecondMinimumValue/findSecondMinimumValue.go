package findSecondMinimumValue

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

const intMax = 1<<63 - 1

func FindSecondMinimumValue(root *TreeNode) int {
	res := intMax

	helper(root, root.Val, &res)

	if res == intMax {
		return -1
	}

	return res
}

func helper(root *TreeNode, lo int, hi *int) {
	if root == nil {
		return
	}

	if lo < root.Val && root.Val < *hi {
		*hi = root.Val
	}

	helper(root.Left, lo, hi)
	helper(root.Right, lo, hi)
}
