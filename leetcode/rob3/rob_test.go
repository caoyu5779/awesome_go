package rob3

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestRob(t *testing.T) {
	t.Run("test rob", func(t *testing.T) {
		pre := []int{2, 3, 3, 3, 1}
		in := []int{3, 2, 3, 3, 1}

		got := Rob(tool.PreIn2Tree(pre, in))
		want := 7

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
