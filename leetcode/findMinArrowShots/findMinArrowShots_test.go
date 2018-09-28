package findMinArrowShots

import (
	"reflect"
	"testing"
)

func TestFindMinArrowShots(t *testing.T) {
	t.Run("test find min arrow shot", func(t *testing.T) {
		nums := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}

		got := FindMinArrowShots(nums)
		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}

	})
}
