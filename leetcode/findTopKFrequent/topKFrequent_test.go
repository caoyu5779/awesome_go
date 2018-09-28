package findTopKFrequent

import (
	"reflect"
	"testing"
)

func TestFindTopKFrequent(t *testing.T) {
	t.Run("test find top k frequent", func(t *testing.T) {
		nums := []int{1, 1, 1, 2, 2, 3}
		k := 2

		got := FindTopKFrequent(nums, k)
		want := []int{1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}

	})
}
