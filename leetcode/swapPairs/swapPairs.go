package swapPairs

import "selfLearning/leetcode/tool"

func SwapPairs(head *tool.ListNode) *tool.ListNode  {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = SwapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}
