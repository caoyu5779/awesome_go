package MinStack

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	t.Run("test min stack", func(t *testing.T) {
		s := Constructor()

		s.Push(-2)
		s.Push(0)
		s.Push(-3)

		got := s.GetMin()
		want := -3

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}

	})
}
