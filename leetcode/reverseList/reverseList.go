package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode

	//head 是要被逆转的节点
	for head != nil {
		//temp 指向 head.Next
		temp := head.Next
		//head.成为已经逆转节点的新head
		head.Next = prev
		//prev 成为 所有已经被逆转节点的head
		prev = head
		// head指向下一个被逆转的节点
		head = temp
	}

	return prev
}
