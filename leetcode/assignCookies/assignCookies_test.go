package assignCookies

import (
	"reflect"
	"testing"
)

func TestAssignCookies(t *testing.T) {
	t.Run("test assign cookies", func(t *testing.T) {
		nums := []int{1, 2}
		s := []int{1, 2, 3}

		got := AssignCookies(nums, s)

		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
