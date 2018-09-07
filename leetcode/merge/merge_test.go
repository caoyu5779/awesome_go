package merge

import (
	"testing"
	"reflect"
)

func TestMerge(t *testing.T) {
	t.Run("test merge" , func(t *testing.T) {
		num1 := []int {1,2,3,0,0,0}
		num2 := []int {2,5,6}

		got := Merge(num1, 3, num2, 3)
		want := []int {1,2,2,3,5,6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}

	})
}
