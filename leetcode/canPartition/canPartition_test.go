package canPartition

import (
	"testing"
	"reflect"
)

func TestCanPartition(t *testing.T) {
	t.Run("test can partition", func(t *testing.T) {
		nums := []int {1,5,11,5}

		got := CanPartition(nums)

		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
