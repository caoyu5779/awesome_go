package Fib

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	t.Run("base Condition", func(t *testing.T) {
		got := fib(3)
		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})

}
