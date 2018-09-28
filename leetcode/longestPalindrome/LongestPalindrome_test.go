package longestPalindrome

import (
	"testing"
	"reflect"
)

func TestLongestPalindrome(t *testing.T) {
	t.Run("test longest palindrome", func(t *testing.T) {
		got := LongestPalindrome("babad")
		want := "abbbbbbbbba"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v , want : %v", got, want)
		}
	})
}
