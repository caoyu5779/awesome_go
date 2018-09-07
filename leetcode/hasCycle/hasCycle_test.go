package hasCycle

import (
	"testing"
	"reflect"
)

func TestHasCycle(t *testing.T) {
	t.Run("test has cycle", func(t *testing.T) {
		node1 := Node{
			nil,
			1,
		}

		node2 := Node{
			node1.Next,
			2,
		}

		node3 := Node{
			node2.Next,
			3,
		}

		node4 := Node{
			node3.Next,
			1,
		}

		got := HasCycle(node4)
		want := true

		if !reflect.DeepEqual(got, want){
			t.Errorf("got : %v , want : %v", got, want)
		}

	})
}
