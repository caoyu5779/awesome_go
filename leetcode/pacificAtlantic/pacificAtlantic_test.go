package pacificAtlantic

import (
	"testing"
	"reflect"
)

func TestPacificAtlantic(t *testing.T) {
	t.Run("test pacific atlantic", func(t *testing.T) {
		grid := [][]int{
			{1, 2, 2, 3, 5},
			{3, 2, 3, 4, 4},
			{2, 4, 5, 3, 1},
			{6, 7, 1, 4, 5},
			{5, 1, 1, 2, 4},
		}

		got := PacificAtlantic(grid)
		want := [][]int{
			{0, 4},
			{1, 3},
			{1, 4},
			{2, 2},
			{3, 0},
			{3, 1},
			{4, 0},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}