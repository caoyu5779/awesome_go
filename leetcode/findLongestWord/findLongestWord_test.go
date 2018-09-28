package findLongestWord

import (
	"reflect"
	"testing"
)

func TestFindLongestWord(t *testing.T) {
	s := "abpcplea"
	d := []string{"ale", "apple", "monkey", "plea"}

	got := FindLongestWord(s, d)
	want := "apple"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got : %v, want : %v", got, want)
	}
}
