package dailyTemperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	t.Run("test daily temperatures", func(t *testing.T) {
		nums := []int{73, 74, 75, 71, 69, 72, 76, 73}

		got := DailyTemperatures(nums)

		want := []int{1, 1, 4, 2, 1, 1, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
