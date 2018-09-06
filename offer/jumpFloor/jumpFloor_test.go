package jumpFloor

import (
	"testing"
	"reflect"
)

func TestJumpFloor(t *testing.T) {
	t.Run("jumpfloor", func(t *testing.T) {
		got := JumpFloor(3)
		want := 3

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v , want : %v", got, want)
		}
	})
}
