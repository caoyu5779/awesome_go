package isSubtree

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

func IsSubtree(s *TreeNode, t *TreeNode) bool {
	if isSame(s, t) {
		return true
	}

	if s == nil {
		return false
	}

	return IsSubtree(s.Left, t) || IsSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}

	if t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}
