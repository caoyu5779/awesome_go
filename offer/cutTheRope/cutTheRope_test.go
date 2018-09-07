package cutTheRope

import (
	"testing"
	"reflect"
)

func TestCutTheRope(t *testing.T)  {
	t.Run("cut the rope", func(t *testing.T) {
		got := cutTheRope(10)
		want := 36

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
