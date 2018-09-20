package maxDepth

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestMaxDepth(t *testing.T) {
	t.Run("test max depth ", func(t *testing.T) {
		 nums := []int {3,9,20, 15,10,7}

		 got := MaxDepth(tool.Ints2TreeNode(nums))
		 want := 3

		 if !reflect.DeepEqual(got, want) {
		 	t.Errorf("got : %v ; want : %v", got, want)
		 }
	})
}
