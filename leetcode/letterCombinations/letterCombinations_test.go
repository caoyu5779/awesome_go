package letterCombinations

import (
	"testing"
	"reflect"
)

func TestLetterCombinations(t *testing.T) {
	t.Run("test letter combination", func(t *testing.T) {
		got := LetterCombinations("23")

		want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
		
	})
}
