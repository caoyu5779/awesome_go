package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

func (node Node) Print() {
	fmt.Print(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil")
		return
	}
	node.Value = value
}

//
//func main()  {
//	//要改变内容 使用指针接收者。 结构过大考虑使用指针接收者
//	//一致性，如果有指针接收者，最好都是指针接收者
//
//	root := Node{Value:3}
//	root.Left = &Node{}
//	root.Right = &Node{5,nil,nil}
//	root.Right.Left = new (Node)
//	root.Left.Right = createNode(2)
//	root.Right.Left.SetValue(4)
//
//	//root.left.right = createNode(2)
//	//
//	//nodes := []Node{
//	//	{value:3},
//	//	{},
//	//	{6,nil,&root},
//	//}
//	//fmt.Println(nodes)
//	//
//	//root.print()
//	//fmt.Println()
//	//
//	//var pRoot *Node
//	//pRoot.setValue(100)
//	//pRoot = &root
//	//pRoot.setValue(300)
//	//pRoot.print()
//
//	root.Traverse()
//}
