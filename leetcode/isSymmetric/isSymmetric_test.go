package isSymmetric

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	t.Run("test is symmetric", func(t *testing.T) {
		pre := []int{1, 2, 3, 4, 2, 4, 3}
		in := []int{3, 2, 4, 1, 4, 2, 3}

		got := IsSymmetric(tool.PreIn2Tree(pre, in))
		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
