package reverseList

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	reversedNums := []int{5, 4, 3, 2, 1}

	got := ReverseList(S2l(nums))

	want := S2l(reversedNums)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got : %v ; want : %v", got, want)
	}
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
		Val: nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}
