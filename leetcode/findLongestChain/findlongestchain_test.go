package findLongestChain

import (
	"reflect"
	"testing"
)

func TestFindLongestChain(t *testing.T) {
	t.Run("test fina longest chain", func(t *testing.T) {
		grid := [][]int{
			{3, 4},
			{2, 3},
			{1, 2},
		}

		got := FindLongestChain(grid)
		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
