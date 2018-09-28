package numberOfArithmeticSlices

import (
	"reflect"
	"testing"
)

func TestNumberOfArithmeticSlice(t *testing.T) {
	a := []int{1, 2, 3, 4}

	got := NumberOfArithmeticSlice(a)
	//
	//want := [][]int {
	//	{1,2,3},
	//	{2,3,4},
	//	{1,2,3,4},
	//}
	want := 3

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got : %v ; want : %v", got, want)
	}

}
