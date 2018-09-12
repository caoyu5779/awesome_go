package searchRange

import (
	"testing"
	"reflect"
)

func TestSearchRange(t *testing.T) {
	t.Run("test search range", func(t *testing.T) {
		nums := []int {5,7,7,8,8,10}
		target := 8

		got := SearchRange(nums, target)
		want := []int {3,4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
		
	})
}
