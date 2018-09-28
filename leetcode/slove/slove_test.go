package slove

import (
	"reflect"
	"testing"
)

func TestSlove(t *testing.T) {
	t.Run("test slove", func(t *testing.T) {
		grid := [][]byte{
			[]byte("OXXXXXOO"),
			[]byte("OOOXXXXO"),
			[]byte("XXXXOOOO"),
			[]byte("XOXOOXXX"),
			[]byte("OXOXXXOO"),
			[]byte("OXXOOXXO"),
			[]byte("OXOXXXOO"),
			[]byte("OXXXXOXX"),
		}

		got := Solve(grid)

		want := [][]byte{
			[]byte("OXXXXXOO"),
			[]byte("OOOXXXXO"),
			[]byte("XXXXOOOO"),
			[]byte("XXXOOXXX"),
			[]byte("OXXXXXOO"),
			[]byte("OXXXXXXO"),
			[]byte("OXXXXXOO"),
			[]byte("OXXXXOXX"),
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
