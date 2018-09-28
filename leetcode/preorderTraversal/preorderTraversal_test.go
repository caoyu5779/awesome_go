package preorderTraversal

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	t.Run("test pre order traversal", func(t *testing.T) {
		pre := []int{1, 2, 4, 5, 3, 6, 7}
		in := []int{4, 2, 5, 1, 6, 3, 7}

		got := PreorderTraversal(tool.PreIn2Tree(pre, in))
		want := []int{4, 5, 2, 6, 7, 3, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
