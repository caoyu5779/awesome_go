package exist

import (
	"testing"
	"reflect"
)

func TestExist(t *testing.T) {
	t.Run("test exist", func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}

		target := "SEE"

		got := Exist(board, target)

		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want )
		}
	})
}