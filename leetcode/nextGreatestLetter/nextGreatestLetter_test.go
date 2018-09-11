package nextGreatestLetter

import (
	"testing"
	"reflect"
)

func TestNextGreatestLetter(t *testing.T) {
	t.Run("test next greatest letter", func(t *testing.T) {
		letters := []byte {'c', 'f', 'j'}
		target := byte('c')

		got := string(NextGreatestLetter(letters, target))
		want := "f"

		if ! reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
		
	})
}
