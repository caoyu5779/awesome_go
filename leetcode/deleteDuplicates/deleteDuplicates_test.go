package deleteDuplicates

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestDeleteDuplicates(t *testing.T) {
	t.Run("test delete duplicates", func(t *testing.T) {
		nums := []int {1,1,2,2,3}
		deletedNums := []int {1,2,3}

		got := DeleteDuplicates(tool.S2l(nums))

		want := tool.S2l(deletedNums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
