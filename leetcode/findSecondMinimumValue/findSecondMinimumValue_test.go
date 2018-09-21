package findSecondMinimumValue

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestFindSecondMinimumValue(t *testing.T) {
	t.Run("test find second minimun value", func(t *testing.T) {
		pre := []int{2,2,5,5,7}
		in := []int{2,2,5,5,7}

		got := FindSecondMinimumValue(tool.PreIn2Tree(pre,in))
		want := 5

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v" , got , want)
		}
	})
}