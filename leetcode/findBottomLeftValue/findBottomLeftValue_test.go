package findBottomLeftValue

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestFindBottomLeftValue(t *testing.T) {
	t.Run("test find bottom left value", func(t *testing.T) {
		pre := []int{1, 2, 4, 3, 5, 7, 6}
		in := []int{4, 2, 1, 7, 5, 3, 6}

		got := FindBottomLeftValue(tool.PreIn2Tree(pre, in))

		want := 7

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
