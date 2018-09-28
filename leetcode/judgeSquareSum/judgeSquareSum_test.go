package judgeSquareSum

import (
	"reflect"
	"testing"
)

func TestJudgeSquareSum(t *testing.T) {
	t.Run("judge square sum", func(t *testing.T) {
		got := JudgeSquareSum(5)
		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}

	})

}
