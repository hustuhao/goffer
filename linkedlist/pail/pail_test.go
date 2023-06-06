package pail

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(vs []int) *ListNode {
	var head *ListNode
	var p *ListNode
	for _, v := range vs {
		if head == nil {
			head = new(ListNode)
			head.Val = v
			p = head
			continue
		}

		node := new(ListNode)
		node.Val = v
		p.Next = node
		p = p.Next
	}
	return head
}

/**
 *
 * @param head ListNode类 the head
 * @return bool布尔型
 */
func isPail(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := new(ListNode)
	slow := new(ListNode)
	var pre *ListNode
	fast = head
	slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		// 反转slow前面的链表
		tmp := slow.Next
		slow.Next = pre
		pre = slow
		slow = tmp
	}
	c1 := pre
	c2 := slow
	if c1 == nil { // 边界条件处理
		return true
	}
	if fast != nil { // 奇数个处理
		c2 = slow.Next
	}
	return isSame(c1, c2)
}

func isSame(h1, h2 *ListNode) (b bool) {
	for h1 != nil && h2 != nil {
		if h1.Val != h2.Val {
			b = false
			return
		}
		h1 = h1.Next
		h2 = h2.Next
	}
	if h1 == nil && h2 == nil {
		b = true
	}
	return
}

func TestPail(t *testing.T) {
	h1 := NewListNode([]int{})
	h2 := NewListNode([]int{1})
	h3 := NewListNode([]int{1, 1})
	h4 := NewListNode([]int{1, 2})
	h5 := NewListNode([]int{1, 2, 1})
	h6 := NewListNode([]int{1, 2, 3})
	h7 := NewListNode([]int{1, 2, 2, 1})
	h8 := NewListNode([]int{1, 2, 3, 1})

	type arg struct {
		wants []bool
		hs    []*ListNode
	}

	as := arg{
		wants: []bool{false, true, true, false, true, false, true, false},
		hs:    []*ListNode{h1, h2, h3, h4, h5, h6, h7, h8},
	}

	for i, v := range as.hs {
		got := isPail(v)
		if got != as.wants[i] {
			t.Fatalf("case %d error", i+1)
		}
	}
}
