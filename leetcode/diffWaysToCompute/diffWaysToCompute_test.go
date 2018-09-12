package diffWaysToCompute

import (
	"testing"
	"reflect"
)

func TestDiffWaysToCompute(t *testing.T) {
	t.Run("test diff ways to compute", func(t *testing.T) {
		s := "2+2+2+2"

		got := DiffWaysToCompute(s)
		want := []int{8,8,8,8,8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v; want : %v", got,want)
		}
	})
}
