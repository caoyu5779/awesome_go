package splitListToParts

import "selfLearning/leetcode/tool"

type ListNode = tool.ListNode

func SplitListToParts(root *ListNode, k int) []*ListNode{
	size := length(root)

	reminder := size % k

	res := make([]*ListNode, k)

	i := 0

	for {
		res[i] = root
		i ++

		if i == k {
			break
		}

		leng := size / k

		if reminder > 0 {
			leng++
			reminder--
		}

		for leng > 1 && root != nil {
			root = root.Next
			leng--
		}

		if root == nil {
			break
		}

		root.Next, root = nil ,root.Next
	}

	return res
}

func length(root *ListNode) int{
	res := 0

	for root != nil {
		res++
		root = root.Next
	}

	return res
}