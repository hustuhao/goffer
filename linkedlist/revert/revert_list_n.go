package revert

import "goffer/common/listnode"

// 反转链表前N个节点
// 链表：1 -> 2 -> 3 -> 4 -> 5, n = 3
// 反转后的链表：3 -> 2 -> 1 -> 4 -> 5
func revertListN(head *listnode.ListNode, n int) *listnode.ListNode {
	var pre *listnode.ListNode
	var cur *listnode.ListNode = head
	var next *listnode.ListNode = cur.Next
	for i := 0; i < n; i++ {
		if cur == nil {
			break
		}
		cur.Next = pre
		pre = cur
		cur = next
		if i == n-1 {
			head.Next = cur
		}
	}
	return pre
}

// https://leetcode.cn/problems/reverse-nodes-in-k-group/
// K 个一组翻转链表
