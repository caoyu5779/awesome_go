package getMinimumDifference

import "selfLearning/leetcode/tool"

type TreeNode = tool.TreeNode

type state struct {
	minDiff, previous int
}

func GetMinimumDifference(root *TreeNode) int {
	st := state{1024, 1024}
	search(root, &st)
	return st.minDiff
}

func search(root *TreeNode, st *state) {
	if root == nil {
		return
	}

	search(root.Left, st)
	newDiff := diff(st.previous, root.Val)

	if st.minDiff > newDiff {
		st.minDiff = newDiff
	}

	st.previous = root.Val

	search(root.Right, st)
}

func diff(i, j int) int {
	if i > j {
		return i - j
	}

	return j - i
}
