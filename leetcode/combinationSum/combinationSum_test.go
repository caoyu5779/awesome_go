package combinationSum

import (
	"testing"
	"reflect"
)

func TestCombinationSum(t *testing.T) {
	t.Run("test combination sum" , func(t *testing.T) {
		put := []int{2, 3, 6, 7}

		got := CombinationSum(put, 7)

		want := [][]int {
			{2,2,3},
			{7},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
