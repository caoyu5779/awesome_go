package removeNthFromEnd

import (
	"fmt"
	"selfLearning/leetcode/tool"
)

func RemoveNthFromEnd(head *tool.ListNode, n int) *tool.ListNode {
	d, headIsNthFromEnd := getParentNode(head, n)

	if headIsNthFromEnd {
		return head.Next
	}

	d.Next = d.Next.Next
	return head
}

func getParentNode(head *tool.ListNode, n int) (parent *tool.ListNode, headIsNthFromEnd bool) {
	parent = head

	for head != nil {
		if n < 0 {
			parent = parent.Next
			fmt.Println(head)
		}

		n--
		fmt.Println(n)
		head = head.Next
	}
	headIsNthFromEnd = n == 0

	return
}
