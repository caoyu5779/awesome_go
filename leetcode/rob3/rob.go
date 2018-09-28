package rob3

import (
	"fmt"
	"selfLearning/leetcode/tool"
)

type TreeNode = tool.TreeNode

func Rob(root *TreeNode) int {
	var dfs func(root *TreeNode) (int, int)

	dfs = func(root *TreeNode) (a int, b int) {
		if root == nil {
			return 0, 0
		}

		la, lb := dfs(root.Left)
		fmt.Println(la, lb)
		ra, rb := dfs(root.Right)
		fmt.Println(ra, rb)
		fmt.Println()

		a = root.Val + lb + rb
		b = max(la, lb) + max(ra, rb)

		return a, b
	}

	a, b := dfs(root)
	return max(a, b)
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
