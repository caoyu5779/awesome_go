package permute

import (
	"testing"
	"reflect"
)

func TestPermute(t *testing.T) {
	t.Run("test permute", func(t *testing.T) {
		nums := []int{1,2,3}

		got := Permute(nums)

		want := [][]int {
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}

	})
}