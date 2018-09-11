package binarySearch

import (
	"testing"
	"reflect"
)

func TestBinarySearch(t *testing.T) {
	t.Run("test binary search ", func(t *testing.T) {
		nums := []int {1,3,2,4,5,6,9,8,2}
		key := 3

		got := BinarySearch(nums, key)
		want := 1

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}

	})
}
