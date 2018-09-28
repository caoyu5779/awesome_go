package nextGreaterElements

import (
	"reflect"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	t.Run("test next greater elements", func(t *testing.T) {
		nums := []int{1, 2, 1}
		got := NextGreaterElements(nums)

		want := []int{2, -1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
