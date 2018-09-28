package isValid

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	t.Run("test is valid", func(t *testing.T) {
		str := "()[]{}"

		got := IsValid(str)

		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
