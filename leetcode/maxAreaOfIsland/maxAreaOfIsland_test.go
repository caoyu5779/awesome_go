package maxAreaOfIsland

import (
	"testing"
	"reflect"
)

func TestMaxAreaOfIsland(t *testing.T) {
	t.Run("test max area of island", func(t *testing.T) {
		grid := [][]int {
			{1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 0, 1, 1},
			{0, 0, 0, 1, 1},
		}

		got := MaxAreaOfIsland(grid)
		want := 4

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
