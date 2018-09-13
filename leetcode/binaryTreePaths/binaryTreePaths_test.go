package binaryTreePaths

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestBinaryTreePaths(t *testing.T) {
	t.Run("test binary tree paths", func(t *testing.T) {
		pre := []int{1,2,5,3}
		in := []int{2,5,1,3}

		got := BinaryTreePaths(tool.PreIn2Tree(pre, in))
		want := []string{"1->2->5","1->3"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v; want : %v", got, want)
		}
	})
}
