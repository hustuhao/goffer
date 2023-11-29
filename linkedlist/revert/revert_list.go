package revert

import (
	"goffer/common/listnode"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 方法1：迭代法
func reverseList1(head *listnode.ListNode) *listnode.ListNode {
	var pre *listnode.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 方法1.5：对方法一的简写优化
func reverseList2(head *listnode.ListNode) *listnode.ListNode {
	var prev *listnode.ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

// 方法2: 递归；递归版本稍微复杂一些，其关键在于反向工作：假设链表的其余部分已经被反转，现在应该如何反转它前面的部分？
/**
 从实现的角度看，递归代码主要包含三个要素。
1. 终止条件:用于决定什么时候由“递”转“归”。
2. 递归调用:对应“递”，函数调用自身，通常输入更小或更简化的参数。
3. 返回结果:对应“归”，将当前递归层级的结果返回至上一层。
*/
func reverseList3(head *listnode.ListNode) *listnode.ListNode {
	if head == nil || head.Next == nil { // 1.终止条件
		return head
	}
	newHead := reverseList3(head.Next) // 2.递归调用：假设链表的其余部分已经被反转。
	head.Next.Next = head
	head.Next = nil
	return newHead // 3.返回结果
}
