package longestUnivaluePath

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func LongestUnivaluePath(root *TreeNode) int {
	maxLen := 0
	helper(root, &maxLen)
	return maxLen
}

func helper(n *TreeNode, maxLen *int) int {
	if n == nil {
		return 0
	}

	l := helper(n.Left, maxLen)
	r := helper(n.Right, maxLen)
	res := 0

	//左侧单边最长距离
	if n.Left != nil && n.Val == n.Left.Val {
		*maxLen = max(*maxLen, l+1)
		res = max(res, l+1)
	}

	//右侧单边最长距离
	if n.Right != nil && n.Val == n.Right.Val {
		*maxLen = max(*maxLen, r+1)
		res = max(res, r+1)
	}
	//通过根节点的最长边
	if n.Left != nil && n.Val == n.Left.Val && n.Right != nil && n.Val == n.Right.Val {
		*maxLen = max(*maxLen, l+r+2)
	}

	return res

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
