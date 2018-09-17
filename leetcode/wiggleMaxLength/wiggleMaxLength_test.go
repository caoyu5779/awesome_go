package wiggleMaxLength

import (
	"testing"
	"reflect"
)

func TestWiggleMaxLength(t *testing.T) {
	t.Run("test wiggle max length", func(t *testing.T) {
		nums := []int {1,7,4,9,2,5}

		got := WiggleMaxLength(nums)

		want := 6

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
