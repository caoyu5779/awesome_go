package diameterOfBinaryTree

import (
	"selfLearning/leetcode/tool"
	"fmt"
)

type TreeNode = tool.TreeNode

func DiameterOfBinaryTree(root *TreeNode) int {
	_, res := helper(root)

	return res
}

func helper(root * TreeNode) (length, diameter int) {
	if root == nil {
		return
	}

	leftLen , leftDia := helper(root.Left)
	rightLen, rightDia := helper(root.Right)

	length = tool.Max(leftLen, rightLen) + 1
	diameter = tool.Max(leftLen+rightLen, tool.Max(leftDia, rightDia))
	fmt.Println(length)
	fmt.Println(diameter)
	fmt.Println()

	return
}
