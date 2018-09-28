package rebuildTraversal

import (
	"reflect"
	"testing"
)

func TestRebuildBinaryTree(t *testing.T) {
	preOrderTraversal := []int{0, 1, 2}
	inOrderTraversal := []int{1, 0, 2}

	got := RebuildBinaryTree(preOrderTraversal, inOrderTraversal)

	left := Node{
		nil,
		nil,
		1,
	}

	right := Node{
		nil,
		nil,
		2,
	}

	root := Node{
		&left,
		&right,
		0,
	}

	want := &root

	if !TreeEqual(got, want) {
		t.Errorf("got : %v want : %v", got, want)
	}
}

func TreeEqual(node1 *Node, node2 *Node) (equality bool) {
	var result1 []int
	var result2 []int

	result1 = node1.PreOrderTraversal(result1)
	result2 = node2.PreOrderTraversal(result2)

	if reflect.DeepEqual(result1, result2) {
		equality = true
	}

	return false
}
