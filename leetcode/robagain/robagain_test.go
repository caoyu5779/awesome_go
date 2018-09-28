package robagain

import (
	"reflect"
	"testing"
)

func TestRobagain(t *testing.T) {
	t.Run("test tob again", func(t *testing.T) {
		nums := []int{2, 3, 2}

		got := Robagain(nums)
		want := 3

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v; want : %v", got, want)
		}
	})
}
