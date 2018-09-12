package numSquares

import (
	"testing"
	"reflect"
)

func TestNumSquares(t *testing.T) {
	t.Run("test num squares", func(t *testing.T) {
		got := NumSquares(12)
		want := 3

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}

	})
}
