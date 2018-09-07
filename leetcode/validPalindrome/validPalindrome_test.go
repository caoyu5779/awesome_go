package validPalindrome

import (
	"testing"
	"reflect"
)

func TestValidPalindrome(t *testing.T) {
	got := ValidPalindrome("abcba")
	want := true

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got : %v, want : %v", got, want)
	}
}
