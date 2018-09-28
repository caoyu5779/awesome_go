package longestUnivaluePath

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestLongestUnivaluePath(t *testing.T) {
	t.Run("test longest univalue path", func(t *testing.T) {
		pre := []int{5, 4, 1, 1, 5, 5}
		in := []int{1, 4, 1, 5, 5, 5}

		got := LongestUnivaluePath(tool.PreIn2Tree(pre, in))

		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
