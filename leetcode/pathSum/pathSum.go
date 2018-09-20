package pathSum

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func PathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	cnt := 0

	var dfs func(*TreeNode, int)

	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum -= node.Val

		if sum == 0 {
			cnt++
		}

		dfs (node.Left, sum)
		dfs (node.Right, sum)
	}

	dfs (root, sum)

	return cnt + PathSum(root.Left, sum) + PathSum(root.Right, sum)
}
