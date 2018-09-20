package diameterOfBinaryTree

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	t.Run("test diameter of binary tree", func(t *testing.T) {
		pre := []int {1,2,4,5,3}
		in := []int {4,2,5,1,3}
		//post := []int {5,2,4,3,1}
		got := DiameterOfBinaryTree(tool.PreIn2Tree(pre, in))

		want := 3

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
