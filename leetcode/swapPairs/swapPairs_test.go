package swapPairs

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestSwapPairs(t *testing.T) {
	t.Run("test swap pairs", func(t *testing.T) {
		nums := []int {1,2,3,4}
		swapNums := []int {2,1,4,3}

		got := SwapPairs(tool.S2l(nums))

		want := tool.S2l(swapNums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
