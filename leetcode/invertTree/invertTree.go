package invertTree

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func InvertTree(root * TreeNode) * TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil){
		return root
	}

	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)

	return root
}
