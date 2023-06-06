package merge_list_v1

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists 合并多条升序链表
// 参数：lists 表示多条升序链表的数组
// 返回值：合并后的升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	l := lists[0]
	for i := 1; i < len(lists); i++ {
		// 将所有链表依次两两合并
		l = mergeTwoLists(l, lists[i])
	}
	return l
}

// mergeTwoLists 合并两条升序链表
// 参数：a, b 分别表示两条升序链表
// 返回值：合并后的升序链表
func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	// 创建一个空链表作为结果
	dump := new(ListNode)
	p := dump
	// 依次比较两条链表的元素，将较小的加入到结果链表中
	for a != nil && b != nil {
		if a.Val < b.Val {
			p.Next = a
			tmp := a.Next
			a.Next = nil
			a = tmp
		} else {
			p.Next = b
			tmp := b.Next
			b.Next = nil
			b = tmp
		}
		p = p.Next
	}

	// 处理剩下的元素
	if a != nil {
		p.Next = a
	}

	if b != nil {
		p.Next = b
	}
	return dump.Next
}
