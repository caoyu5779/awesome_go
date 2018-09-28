package myAtoi

import (
	"testing"
	"reflect"
)

func TestMyAtoi(t *testing.T) {
	t.Run("test my atoi", func(t *testing.T) {
		got := MyAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459")
		want := 2147483647

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}