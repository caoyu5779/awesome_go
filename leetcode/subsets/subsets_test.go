package subsets

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	t.Run("test subsets", func(t *testing.T) {
		para := []int{1, 2, 3}

		got := Subsets(para)

		want := [][]int{
			{},
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
