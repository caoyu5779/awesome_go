package findCircleNum

import (
	"reflect"
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	t.Run("test find circle num", func(t *testing.T) {
		nums := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}

		got := FindCircleNum(nums)

		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
