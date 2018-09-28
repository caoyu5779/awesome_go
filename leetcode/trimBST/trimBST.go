package trimBST

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func TrimBST(root *TreeNode, l int, r int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < l {
		return TrimBST(root.Right, l, r)
	}

	if root.Val > r {
		return TrimBST(root.Left, l, r)
	}

	root.Left = TrimBST(root.Left, l, r)
	root.Right = TrimBST(root.Right, l, r)

	return root
}
