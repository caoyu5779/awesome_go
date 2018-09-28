package findBottomLeftValue

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func FindBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		//先把Right 放进队列 最终能够保证 最后一个是左节点
		if root.Right != nil {
			queue = append(queue, root.Right)
		}

		if root.Left != nil {
			queue = append(queue, root.Left)
		}

	}

	return root.Val
}
