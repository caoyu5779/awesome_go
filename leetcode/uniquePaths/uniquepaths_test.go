package uniquePaths

import (
"testing"
"reflect"
)

func TestUniquePaths(t *testing.T) {
	t.Run("test unique path", func(t *testing.T) {
		got := UniquePaths(7,3)
		want := 28

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got,want)
		}
	})
}
