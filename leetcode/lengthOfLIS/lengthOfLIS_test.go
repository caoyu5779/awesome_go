package lengthOfLIS

import (
	"reflect"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	t.Run("test length of lis", func(t *testing.T) {
		nums := []int{10, 9, 2, 5, 3, 7, 101, 18}

		got := LengthOfLIS(nums)
		want := 4

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
