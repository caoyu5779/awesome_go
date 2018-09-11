package maxProfit

import (
	"testing"
	"reflect"
)

func TestMaxProfit(t *testing.T) {
	t.Run("test max profit", func(t *testing.T) {
		price := []int {7,1,5,3,6,4}

		got := MaxProfit(price)
		want := 7

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v; want : %v", got, want)
		}

	})
}
