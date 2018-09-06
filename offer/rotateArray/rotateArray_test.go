package rotateArray

import (
	"testing"
	"reflect"
)

func TestRotateArray(t *testing.T)  {
	t.Run("testing rotate array", func(t *testing.T) {
		array := []int {1,2,3,4,5,6}
		got := rotateArray(array, 4)
		want := []int{6,1,2,3,4,5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
