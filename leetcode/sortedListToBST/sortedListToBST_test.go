package sortedListToBST

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestSortedListToBST(t *testing.T) {
	t.Run("test sorted list to bst", func(t *testing.T) {
		head := []int {1,2,3,4,5,6,7,8}
		want := []int {1,2,3,4,5,6,7,8}

		got := tool.Tree2Inorder(SortedListToBST(tool.S2l(head)))

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
