package findMin

import (
	"reflect"
	"testing"
)

func TestFindMin(t *testing.T) {
	t.Run("test find min", func(t *testing.T) {
		input := []int{3, 4, 5, 1, 2}

		got := FindMin(input)

		want := 1

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
