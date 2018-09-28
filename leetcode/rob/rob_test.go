package rob

import (
	"reflect"
	"testing"
)

func TestRob(t *testing.T) {
	t.Run("test rob", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}

		got := Rob(nums)
		want := 4

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
}
