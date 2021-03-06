package preorderTraversal

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func PreorderTraversal(root *TreeNode) []int {
	var rightStack []*TreeNode
	var res []int

	for cur := root; cur != nil; {
		res = append(res, cur.Val)

		if cur.Left != nil {
			if cur.Right != nil {
				rightStack = append(rightStack, cur.Right)
			}
			cur = cur.Left
		} else {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				if len(rightStack) == 0 {
					break
				}
				cur = rightStack[len(rightStack)-1]
				rightStack = rightStack[:len(rightStack)-1]
			}
		}
	}

	return res
}
