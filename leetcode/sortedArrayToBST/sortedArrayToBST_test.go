package sortedArrayToBST

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestSortedArrayToBST(t *testing.T) {
	t.Run("test sorted array to bst", func(t *testing.T) {
		nums := []int {1,2,3,4,5,6,7}
		pre := []int{4, 2, 1, 3, 6, 5, 7}

		got := tool.Tree2PreOrder(SortedArrayToBST(nums))
		want := pre

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
