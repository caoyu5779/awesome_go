package tool

type ListNode struct {
	Val int
	Next *ListNode
}

func L2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func S2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val : nums[0],
	}

	temp := res

	for i:=1; i < len(nums) ; i ++ {
		temp.Next = &ListNode{
			Val:nums[i],
		}
		temp = temp.Next
	}

	return res
}


