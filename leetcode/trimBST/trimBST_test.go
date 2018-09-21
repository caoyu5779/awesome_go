package trimBST

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestTrimBST(t *testing.T) {
	t.Run("test trim bst", func(t *testing.T) {
		in := []int{0,1,2}
		pre := []int{1,0,2}

		post := []int{2,1}
		in1 := []int{1,2}

		got := TrimBST(tool.PreIn2Tree(pre, in), 1, 2)
		want := tool.InPost2Tree(in1, post)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}