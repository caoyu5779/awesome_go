package minPathSum

import (
	"testing"
	"reflect"
)

func TestMinPathSum(t *testing.T) {
	t.Run("test min path sum", func(t *testing.T) {
		grid := [][]int {
			{1,3,1},
			{1,5,1},
			{4,2,1},
		}

		got := MinPathSum(grid)
		want := 7

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}

	})
}
