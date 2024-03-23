package ipqueue

// 优先队列（Priority Queue）：一种特殊的队列，等同于大顶堆。
// 在优先队列中，元素被赋予优先级，当访问队列元素时，具有最高优先级的元素最先删除。

// Item 队列中的元素
type Item[T any] struct {
	Value    T     // 元素值
	Priority int64 // 优先级
}

// PriorityQueue 优先队列，使用切片存储元素，并实现了 heap.Interface 接口；可以不使用 heap 的相关操作
type PriorityQueue[T any] []*Item[T]

// 获取队列中的第N个元素
func (pq *PriorityQueue[T]) Get(index int) *Item[T] {
	return (*pq)[index]
}

// Maintains the heap property after inserting an item
func (pq *PriorityQueue[T]) heapifyUp(index int) {
	for (*pq)[index].Priority < (*pq)[(index-1)/2].Priority {
		pq.Swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// Maintains the heap property after removing an item
func (pq *PriorityQueue[T]) heapifyDown(index int) {
	lastIndex := len(*pq) - 1
	l, r := 2*index+1, 2*index+2
	for l <= lastIndex {
		if l == lastIndex { // Only the left child exists
			if (*pq)[l].Priority < (*pq)[index].Priority {
				pq.Swap(l, index)
			}
			return
		} else if (*pq)[l].Priority < (*pq)[r].Priority && (*pq)[l].Priority < (*pq)[index].Priority {
			// Left child is smaller
			pq.Swap(l, index)
			index = l
		} else if (*pq)[r].Priority < (*pq)[index].Priority {
			// Right child is smaller
			pq.Swap(r, index)
			index = r
		} else {
			return // Heap property is maintained
		}
		l, r = 2*index+1, 2*index+2
	}
}

func (pq *PriorityQueue[T]) Len() int           { return len(*pq) }
func (pq *PriorityQueue[T]) Less(i, j int) bool { return (*pq)[i].Priority < (*pq)[j].Priority }
func (pq *PriorityQueue[T]) Swap(i, j int)      { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }

// Offer 添加元素到队列中，并调用 heapifyUp()
func (pq *PriorityQueue[T]) Offer(item *Item[T]) {
	*pq = append(*pq, item)
	pq.heapifyUp(len(*pq) - 1)
}

// Poll 移除最高优先级的元素，并调用 heapifyDown() 方法保持堆的特性。
func (pq *PriorityQueue[T]) Poll() *Item[T] {
	top := (*pq)[0]
	(*pq)[0] = (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	pq.heapifyDown(0)
	return top
}

// NewQueue returns a new priority queue with the given items.
func NewQueue[T any](items ...*Item[T]) *PriorityQueue[T] {
	pq := new(PriorityQueue[T])
	for _, item := range items {
		pq.Offer(item)
	}
	return pq
}
