package maxSubArray

import (
	"testing"
	"reflect"
)

func TestMaxSubArray(t *testing.T) {
	t.Run("test max sub array", func(t *testing.T) {
		nums := []int {-2,1,-3,4,-1,2,1,-5,4}

		got := MaxSubArray(nums)

		want := 6

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
