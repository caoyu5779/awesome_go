package sumOfLeftLeaves

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestSumOfLeftLeaves(t *testing.T) {
	t.Run("test sum of left leaves", func(t *testing.T) {
		pre := []int {3,9,20,15,7}
		in := []int {9,3,15,20,7}

		got := SumOfLeftLeaves(tool.PreIn2Tree(pre, in))

		want := 24

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
