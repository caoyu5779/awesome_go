package hasPathSum

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestHasPathSum(t *testing.T) {
	t.Run("test has path sum", func(t *testing.T) {
		pre := []int {5,4,11,7,2,8,13,4,1}
		in := []int {7,11,2,4,5,13,8,4,1}

		got := HasPathSum(tool.PreIn2Tree(pre,in), 22)
		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
