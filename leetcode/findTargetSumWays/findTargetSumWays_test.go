package findTargetSumWays

import (
	"testing"
	"reflect"
)

func TestFindTargetSumWays(t *testing.T) {
	t.Run("test find target sum ways", func(t *testing.T) {
		nums := []int {1,1,1,1,1}

		got := FindTargetSumWays(nums, 3)

		want := 5

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}