package validPalindrome

import (
	"reflect"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	got := ValidPalindrome("abcba")
	want := true

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got : %v, want : %v", got, want)
	}
}
