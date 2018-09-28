package isPalindrome

import "selfLearning/leetcode/tool"

type ListNode = tool.ListNode

func IsPalindrome(head *ListNode) bool {
	nums := make([]int, 0, 64)

	//转换成数组
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	l, r := 0, len(nums)-1
	//从两头开始比较
	for l < r {
		if nums[l] != nums[r] {
			return false
		}
		l++
		r--
	}

	return true
}
