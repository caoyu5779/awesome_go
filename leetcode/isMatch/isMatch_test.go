package isMatch

import (
	"testing"
	"reflect"
)

func TestIsMatch(t *testing.T) {
	t.Run("test is match", func(t *testing.T) {
		got := IsMatch("aa", "a")
		want := false

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}