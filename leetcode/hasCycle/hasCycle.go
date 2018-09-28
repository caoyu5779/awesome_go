package hasCycle

import "fmt"

type Node struct {
	Next  *Node
	Value interface{}
}

func HasCycle(node *Node) bool {
	var first *Node

	visited := make(map[*Node]bool)

	for n := first; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			return true
		}

		visited[n] = true
		fmt.Println(n.Value)
	}

	return false
}
