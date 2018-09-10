package sortColors

import (
	"testing"
	"reflect"
)

func TestSortColors(t *testing.T) {
	t.Run("test color sort", func(t *testing.T) {
		nums := []int {0,1,2,0,1,2}
		got := SortColors(nums)

		want := []int {0,0,1,1,2,2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}

	})
}
