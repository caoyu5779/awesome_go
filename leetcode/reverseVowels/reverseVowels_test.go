package reverseVowels

import (
	"testing"
	"reflect"
)

func TestReverseVowels(t *testing.T) {
	t.Run("reverse vowels", func(t *testing.T) {
		s := "leetcode"
		got := ReverseVowels(s)
		want := "leotcede"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
