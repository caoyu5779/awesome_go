package lengthOfLongestSubstring

import (
	"testing"
	"reflect"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("test length of longest substring", func(t *testing.T) {
		got := LengthOfLongestSubstring("abcabcbb")
		want := 3

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
