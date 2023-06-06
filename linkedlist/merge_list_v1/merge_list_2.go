package merge_list_v1

import "container/heap"

// 优先队列（最小堆 ）

//  IntHeap 使用了指针接收器，使得所有的方法都能使用指针访问 IntHeap 的元素。
// 由于使用指针接收器会增加一些额外的内存开销，因此这个版本的代码可能会比原来的版本稍微慢一些，但是差别并不大。

type IntHeap []*ListNode

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *IntHeap) Swap(i, j int) {
	tmp := (*h)[i]
	(*h)[i] = (*h)[j]
	(*h)[j] = tmp
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *IntHeap) Pop() any {
	n := h.Len() - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func mergeKListsV2(lists []*ListNode) *ListNode {
	h := new(IntHeap)
	heap.Init(h)
	for _, head := range lists {
		if head == nil {
			continue
		}
		heap.Push(h, head)
	}

	dump := new(ListNode)
	p := dump
	for h.Len() > 0 {
		n := heap.Pop(h).(*ListNode)
		p.Next = n
		tmp := n.Next
		if n.Next != nil {
			n.Next = nil
			heap.Push(h, tmp)
		}
		p = p.Next
	}
	return dump.Next
}
