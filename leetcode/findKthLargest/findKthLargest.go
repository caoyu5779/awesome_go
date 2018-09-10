package findKthLargest

import "container/heap"

func FindKthLargest(nums []int, k int) int {
	temp := highHeap(nums)

	h := &temp

	heap.Init(h)

	if k==1 {
		return (*h)[0]
	}

	for i:= 1; i<k; i++ {
		heap.Remove(h, 0)
	}

	return (*h)[0]
}

type highHeap []int

func (h highHeap) Len() int  {
	return len(h)
}

func (h highHeap) Less(i, j int) bool{
	return h[i] > h[j]
}

func (h highHeap) Swap (i,j int) {
	h[i],h[j] = h[j], h[i]
}

func (h *highHeap) Push (x interface{})  {
	//push 使用 指针 是因为push 增加了h的长度
	*h = append(*h , x.(int))
}

func (h *highHeap) Pop() interface{} {
	// pop使用指针是因为pop 减短了 h的长度
	res := (*h)[len(*h) - 1]
	*h = (*h)[0: len(*h) - 1]

	return res
}



