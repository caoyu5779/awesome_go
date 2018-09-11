package reconstructQueue

import (
	"testing"
	"reflect"
)

func TestReconstructQueue(t *testing.T) {
	t.Run("test reconstruct queue", func(t *testing.T) {
		people := [][]int {{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}

		got := ReconstructQueue(people)

		want := [][]int {{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v , want : %v", got, want)
		}

	})
}
