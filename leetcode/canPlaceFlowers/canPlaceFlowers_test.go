package canPlaceFlowers

import (
	"reflect"
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	t.Run("test can place flowers", func(t *testing.T) {
		flowers := []int{1, 0, 0, 0, 1, 0}
		n := 1
		got := CanPlaceFlowers(flowers, n)
		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}

	})
}
