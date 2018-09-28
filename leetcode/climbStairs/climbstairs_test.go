package climbStairs

import (
	"reflect"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	t.Run("test climb stairs", func(t *testing.T) {
		got := ClimbStairs(3)

		want := 3

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
