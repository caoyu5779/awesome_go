package splitListToParts

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestSplitListToParts(t *testing.T) {
	t.Run("test split list to parts", func(t *testing.T) {
		root := []int {1,2,3}

		got := SplitListToParts(tool.S2l(root), 5)
		want := []*ListNode {
			tool.S2l([]int{1}),
			tool.S2l([]int{2}),
			tool.S2l([]int{3}),
			nil,
			nil,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got  : %v; want : %v", got, want)
		}
	})
}
