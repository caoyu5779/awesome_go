package convert

import (
	"testing"
	"reflect"
)

func TestConvert(t *testing.T) {
	t.Run("test convert", func(t *testing.T) {
		got := Convert("PAYPALISHIRING", 3)
		want := "PAHNAPLSIIGYIR"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})

}
