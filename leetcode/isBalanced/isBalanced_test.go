package isBalanced

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	t.Run("test is balanced", func(t *testing.T) {
		nums := []int{4, 2, 1, 3, 6, 5, 7}
		post := []int{1, 2, 3, 4, 5, 6, 7}

		got := IsBalanced(tool.InPost2Tree(nums, post))

		want := false

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
