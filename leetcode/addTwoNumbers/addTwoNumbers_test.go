package addTwoNumbers

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("test add two numbers", func(t *testing.T) {
		nums := []int{7, 2, 4, 3}
		nums_two := []int{5, 6, 4}
		addNums := []int{7, 8, 0, 7}

		got := AddTwoNumbers(tool.S2l(nums), tool.S2l(nums_two))

		want := tool.S2l(addNums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
