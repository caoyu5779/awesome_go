package minDepth

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestMinDepth(t *testing.T) {
	t.Run("test min depth", func(t *testing.T) {
		pre := []int{3, 9, 20, 15, 7}
		in := []int{9, 3, 15, 20, 7}

		got := MinDepth(tool.PreIn2Tree(pre, in))

		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
