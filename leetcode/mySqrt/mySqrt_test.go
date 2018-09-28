package mySqrt

import (
	"reflect"
	"testing"
)

func TestMySqrt(t *testing.T) {
	t.Run("test my sqrt", func(t *testing.T) {
		got := MySqrt(4)
		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}

	})
}
