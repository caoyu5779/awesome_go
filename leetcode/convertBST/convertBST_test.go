package convertBST

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestConvertBST(t *testing.T) {
	t.Run("test convert bst", func(t *testing.T) {
		in := []int {2,5,13}
		pre := []int {5,2,13}

		in1 := []int {20,18,13}
		pre1 := []int {18,20,13}

		got := ConvertBST(tool.PreIn2Tree(pre, in))

		want := tool.PreIn2Tree(pre1, in1)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
