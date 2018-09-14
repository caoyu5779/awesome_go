package partition

import (
	"testing"
	"reflect"
)

func TestPartition(t *testing.T) {
	t.Run("test partition", func(t *testing.T) {
		got := Partition("aab")
		want := [][]string {
			{"a", "a", "b"},
			{"aa", "b"},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
