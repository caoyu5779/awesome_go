package permutationNum

import (
	"testing"
	"reflect"
)

func TestPermutationNum(t *testing.T) {
	t.Run("run test", func(t *testing.T) {
		got := PermutationNum(1)
		want := 0

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
