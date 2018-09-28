package isPalindrome

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	t.Run("'test is palindrome", func(t *testing.T) {
		nums := []int{1, 2, 3, 2, 1}

		got := IsPalindrome(tool.S2l(nums))
		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
