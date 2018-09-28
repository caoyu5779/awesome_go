package mergeTrees

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	t.Run("test merge trees", func(t *testing.T) {
		pre1 := []int{1, 3, 5, 2}

		in1 := []int{5, 3, 1, 2}

		pre2 := []int{20, 10, 40, 30, 70}
		in2 := []int{10, 40, 20, 30, 70}
		nums := []int{5, 40, 13, 70, 32, 21}

		got := MergeTrees(tool.PreIn2Tree(pre1, in1), tool.PreIn2Tree(pre2, in2))
		want := tool.Ints2TreeNode(nums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
