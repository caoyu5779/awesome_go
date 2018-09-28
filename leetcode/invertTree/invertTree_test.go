package invertTree

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestInvertTree(t *testing.T) {
	t.Run("test invert tree", func(t *testing.T) {
		pre := []int{4, 2, 1, 3, 7, 6, 9}
		in := []int{1, 2, 3, 4, 6, 7, 9}
		final := []int{4, 7, 2, 9, 6, 3, 1}

		got := InvertTree(tool.PreIn2Tree(pre, in))
		want := tool.Ints2TreeNode(final)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
