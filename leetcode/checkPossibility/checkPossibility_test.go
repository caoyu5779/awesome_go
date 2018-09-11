package checkPossibility

import (
	"testing"
	"reflect"
)

func TestCheckPossibility(t *testing.T) {
	t.Run("test check possibility", func(t *testing.T) {
		input := []int {4,2,3}

		got := CheckPossibility(input)
		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
