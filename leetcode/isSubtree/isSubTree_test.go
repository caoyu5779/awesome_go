package isSubtree

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestIsSubtree(t *testing.T) {
	t.Run("test is sub tree", func(t *testing.T) {
		pre1 := []int{3, 4, 1, 2, 5}
		in1 := []int{1, 4, 2, 3, 5}

		pre2 := []int{4, 1, 2}
		in2 := []int{1, 4, 2}

		got := IsSubtree(tool.PreIn2Tree(pre1, in1), tool.PreIn2Tree(pre2, in2))

		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
