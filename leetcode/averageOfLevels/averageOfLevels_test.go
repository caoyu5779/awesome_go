package averageOfLevels

import (
	"reflect"
	"selfLearning/leetcode/tool"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	t.Run("test average of levels", func(t *testing.T) {
		in := []int{9, 3, 15, 20, 7}
		pre := []int{3, 9, 20, 15, 7}

		got := AverageOfLevels(tool.PreIn2Tree(pre, in))
		want := []float64{3, 14.5, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
