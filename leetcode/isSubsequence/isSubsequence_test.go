package isSubsequence

import (
	"testing"
	"reflect"
)

func TestIsSubsequence(t *testing.T) {
	t.Run("test is subsequence", func(t *testing.T) {
		s := "abc"
		b := "ahbgdc"

		got := IsSubsequence(s, b)
		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
