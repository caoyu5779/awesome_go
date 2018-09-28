package combine

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	t.Run("test combine", func(t *testing.T) {
		got := Combine(4, 2)

		want := [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}

	})
}
