package convertBST

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func ConvertBST(root *TreeNode) *TreeNode {
	sum := 0
	travel(root, &sum)
	return root
}

func travel(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	travel(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	travel(root.Left, sum)
}
