package numDecodings

import (
	"reflect"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	t.Run("test num decoding", func(t *testing.T) {
		got := NumDecodings("226")
		want := 3

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
