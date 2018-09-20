package oddEvenList

import (
	"selfLearning/leetcode/tool"
	"fmt"
)

type ListNode = tool.ListNode
func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//oe ： odd end 指向 odd group的最后一个node
	oe := head
	//eb : even begin 指向 even group 的第一个node
	eb := head.Next
	// ee : even end 指向 even group 的最后一个node
	ee := eb

	for ee != nil && ee.Next != nil {
		//->2 = ->3
		oe.Next = ee.Next
		fmt.Println(oe.Next)
		//oe = -> 2
		oe = oe.Next
		fmt.Println(oe)
		// -> 3 = -> 3
		ee.Next = oe.Next
		fmt.Println(ee.Next)
		// ee = -> 3
		ee = ee.Next
		fmt.Println(ee)
		// ->3 = -> 2
		oe.Next = eb
		fmt.Println(oe.Next)
		fmt.Println()
	}

	return head
}
