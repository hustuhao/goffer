package listnode

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 用于构建链，返回链表头
func NewList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, val := range vals {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	return dummy.Next
}

func ListNodeToStrList(head *ListNode) string {
	strs := make([]string, 0)
	for head != nil {
		head = head.Next
		strs = append(strs, fmt.Sprintf("%v", head.Val))
	}
	return strings.Join(strs, ",")
}
