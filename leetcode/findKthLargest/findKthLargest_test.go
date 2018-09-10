package findKthLargest

import (
	"testing"
	"reflect"
)

func TestFindKthLargest(t *testing.T) {
	t.Run("test find kth largest", func(t *testing.T) {
		nums := []int {3,3,3,3,3,12,2,312,23,23123123,1231,0,112}
		k := 3

		var got []int
		for i := 1; i <= k; i++ {
			got = append(got, FindKthLargest(nums, i))
		}
		want := []int {23123123,1231,312}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v , want : %v", got, want)
		}

	})
}
