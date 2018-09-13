package numsIslands

import (
	"testing"
	"reflect"
)

func TestNumIslands(t *testing.T) {
	t.Run("test nums islands", func(t *testing.T) {
		nums := [][]byte{
			[]byte("111001"),
			[]byte("010001"),
			[]byte("111111"),
		}

		got := NumIslands(nums)

		want := 1

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
