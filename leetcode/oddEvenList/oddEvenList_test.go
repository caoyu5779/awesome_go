package oddEvenList

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	t.Run("test odd even list", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		oddNums := []int{1, 3, 5, 2, 4}

		got := OddEvenList(tool.S2l(nums))

		want := tool.S2l(oddNums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
