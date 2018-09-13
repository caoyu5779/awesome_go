package binaryTreePaths

import (
	"selfLearning/leetcode/tool"
	"strconv"
)

type TreeNode = tool.TreeNode

func BinaryTreePaths(root *TreeNode) []string  {

	if root == nil{
		return nil
	}

	res := make([]string, 0, 16)

	var dfs func(string, *TreeNode)

	dfs = func(pre string, root *TreeNode) {
		if pre == "" {
			pre = strconv.Itoa(root.Val)
		}	else {
			pre += "->" + strconv.Itoa(root.Val)
		}

		if  root.Left != nil {
			dfs(pre, root.Left)
		}

		if root.Right != nil {
			dfs(pre, root.Right)
		}

		if root.Left == nil && root.Right == nil {
			res = append(res, pre)
		}
	}

	dfs("", root)

	return res
}