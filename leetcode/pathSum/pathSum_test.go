package pathSum

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestPathSum(t *testing.T) {
	t.Run("test path sum", func(t *testing.T) {
		pre := []int{10, 5, 3, 3, -2, 2, 1, -3, 11}
		in := []int{3, 3, -2, 5, 2, 1, 10, -3, 11}

		got := PathSum(tool.PreIn2Tree(pre, in), 8)

		want := 3

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
