package partitionLabels

import (
	"testing"
	"reflect"
)

func TestPartitionLabels(t *testing.T) {
	t.Run("test partition labels", func(t *testing.T) {
		S := "ababcbacadefegdehijhklij"

		got := PartitionLabels(S)
		want := []int {9,7,8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
		
	})
}
