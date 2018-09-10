package nonOverlapIntervals

import (
	"testing"
	"selfLearning/leetcode/tool"
	"reflect"
)

func TestNonOverlapIntervals(t *testing.T) {
	t.Run("test non overlap intervals", func(t *testing.T) {
		nums := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}

		got := NonOverlapIntervals(tool.Intss2IntervalSlice(nums))
		want := 1

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}