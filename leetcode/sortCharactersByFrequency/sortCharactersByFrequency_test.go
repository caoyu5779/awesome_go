package sortCharactersByFrequency

import (
	"reflect"
	"testing"
)

func TestSortCharactersByFrequency(t *testing.T) {
	t.Run("test sort characters by frequency", func(t *testing.T) {
		got := SortCharactersByFrequency("eaac")
		want := "aace"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
