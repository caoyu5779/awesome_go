package tool

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var NULL = -1 <<63

func Ints2TreeNode(ints []int) *TreeNode  {
	n := len(ints)

	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val : ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1

	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val:ints[i]}
			queue = append(queue, node.Left)
		}

		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val:ints[i]}
			queue = append(queue, node.Right)
		}

		i++
	}

	return root
}

func GetTargetNode(root * TreeNode, target int) * TreeNode {
	if root == nil || root.Val == target {
		return root
	}

	res := GetTargetNode(root.Left, target)

	if res != nil {
		return res
	}

	return GetTargetNode(root.Right, target)
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val{
			return i
		}
	}

	msg := fmt.Sprintf("%d no in %v", val, nums)

	panic(msg)
}
//把preorder 和 inorder 转换为二叉树 前序
func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片长度不同")
	}

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val:pre[0],
	}
	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val,in)

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

//中序
func InPost2Tree(in,post []int) *TreeNode  {
	if len(post) != len(in) {
		panic("in post 2 tree 两个切片长度不同")
	}

	if len(in) == 0{
		return nil
	}

	res := &TreeNode{
		Val:post[len(post) - 1],
	}

	if len(in) == 1{
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = InPost2Tree(in[:idx], post[:idx])
	res.Right = InPost2Tree(in[idx+1:], post[idx:len(post) - 1])

	return res
}

//Tree2PreOrder 二叉树转成 preorder 切片
func Tree2PreOrder(root *TreeNode) []int {
	if root == nil{
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int {root.Val}
	}

	res := []int {root.Val}

	res = append(res, Tree2PreOrder(root.Left)...)
	res = append(res, Tree2PreOrder(root.Right)...)

	return res
}

//Tree2Inorder 二叉树转换成 中序切片

func Tree2Inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int {root.Val}
	}

	res := Tree2Inorder(root.Left)

	res = append(res, root.Val)
	res = append(res, Tree2Inorder(root.Right)...)

	return res
}

//Tree2PostOrder 二叉树转换成 后续切片
func Tree2PostOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int {root.Val}
	}

	res := Tree2PostOrder(root.Left)

	res = append(res, Tree2PostOrder(root.Right)...)
	res = append(res, root.Val)

	return res
}

//equal return true if tn == a
func (tn *TreeNode) Equal (a *TreeNode) bool  {
	if tn == nil && a == nil {
		return true
	}

	if tn == nil || a == nil || tn.Val != a.Val {
		return false
	}

	return tn.Left.Equal(a.Left) && tn.Right.Equal(a.Right)
}

//Tree 2 ints 把*TreeNode 按照行还原成[]int
func Tree2ints(tn *TreeNode) []int {
	res := make([]int, 0, 1024)

	queue := []*TreeNode{tn}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0 ; i< size; i ++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, NULL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}

		queue = queue[size:]
	}

	i := len(res)

	for i > 0 && res[i-1] == NULL {
		i--
	}

	return res[:i]
}
