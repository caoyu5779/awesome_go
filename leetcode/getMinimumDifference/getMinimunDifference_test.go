package getMinimumDifference

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestGetMinimumDifference(t *testing.T) {
	t.Run("test get minimum differenct", func(t *testing.T) {
		pre := []int{1,3,2}
		in := []int{1,2,3}

		got := GetMinimumDifference(tool.PreIn2Tree(pre, in))
		want := 1

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}