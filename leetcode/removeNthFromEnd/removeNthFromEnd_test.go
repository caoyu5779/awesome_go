package removeNthFromEnd

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	removedNums := []int{1, 2, 4, 5}

	got := RemoveNthFromEnd(tool.S2l(nums), 3)
	want := tool.S2l(removedNums)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got : %v ; want : %v", got, want)
	}
}
