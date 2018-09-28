package kthSmallest

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	t.Run("test kth smallest", func(t *testing.T) {
		pre := []int{2, 1, 3}
		in := []int{1, 2, 3}

		got := KthSmallest(tool.PreIn2Tree(pre, in), 2)
		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
