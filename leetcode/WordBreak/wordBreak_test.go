package WordBreak

import (
	"reflect"
	"testing"
)

func TestWordBreak(t *testing.T) {
	t.Run("test word break", func(t *testing.T) {
		s := "applepenapple"
		wordDict := []string{"apple", "pen"}

		got := WordBreak(s, wordDict)
		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
