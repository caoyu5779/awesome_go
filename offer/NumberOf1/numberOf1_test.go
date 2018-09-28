package NumberOf1

import (
	"reflect"
	"testing"
)

func TestNumberOf1(t *testing.T) {
	t.Run("number of 1 ", func(t *testing.T) {
		got := numberOf1(10)
		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}

	})
}
