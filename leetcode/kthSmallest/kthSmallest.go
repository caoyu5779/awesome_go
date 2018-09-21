package kthSmallest

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func KthSmallest(root *TreeNode, k int) int {
	leftSize := getSize(root.Left)

	switch  {
	case k <= leftSize :
		return KthSmallest(root.Left, k)
	case leftSize+1 < k:
		return KthSmallest(root.Right, k - leftSize - 1)
	default:
		return root.Val
	}
}

//获取root节点的数量
func getSize(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + getSize(root.Left) + getSize(root.Right)
}