package sumRange

import (
	"testing"
	"reflect"
)

func TestNumArray_SumRange(t *testing.T) {
	var nums = []int{-2, 0, 3, -5, 2, -1}

	na := Constructor(nums)

	got := na.SumRange(2,5)

	want := -1

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got : %v ; want : %v", got, want)
	}
}