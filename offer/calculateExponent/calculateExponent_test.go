package calculateExponent

import (
	"testing"
	"reflect"
)

func TestCalculateExponent(t *testing.T)  {
	t.Run("calculate exponent", func(t *testing.T) {
		got,err := calculateExponent(10, -2)
		want := 0.01

		if err != nil {
			t.Errorf("error : %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}

	})

}
