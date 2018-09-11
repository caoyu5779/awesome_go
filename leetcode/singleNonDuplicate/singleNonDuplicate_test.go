package singleNonDuplicate

import (
	"testing"
	"reflect"
)

func TestSingleNonDuplicate(t *testing.T) {
	t.Run("test single non duplicate ", func(t *testing.T) {
		nums := []int {3, 3, 7, 7, 10, 11, 11}

		got := SingleNonDuplicate(nums)
		want := 10

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v; want : %v", got, want)
		}
		
	})
}
