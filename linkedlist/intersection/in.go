package intersection

import (
	"goffer/common/listnode"
)

/**
 * 找到两个链表的交点。
 *
 * @param headA 第一个链表的头节点
 * @param headB 第二个链表的头节点
 * @return 两个链表的交点，如果不存在交点则返回 nil
 */
func getIntersectionNode1(headA, headB *listnode.ListNode) *listnode.ListNode {
	// 初始化两个指针，分别指向两个链表的头节点
	p1 := headA
	p2 := headB

	// 如果两个指针都为 nil，则说明两个链表没有交点，直接返回 nil
	// 否则，执行下面的循环
	for !(p1 == nil && p2 == nil) {
		// 如果两个指针指向同一个节点，则说明此节点为两个链表的交点，直接返回该节点
		if p1 == p2 {
			return p1
		}
		// 否则，将两个指针分别向前移动一步
		p1 = p1.Next
		p2 = p2.Next
		// 如果两个指针都移动到了链表末尾，说明两个链表没有交点，直接返回 nil
		if p1 == nil && p2 == nil {
			return nil
		}
		// 如果 p1 已经移动到了链表末尾，则将其重置为 headB，从而将其移动到第二个链表的头节点
		if p1 == nil {
			p1 = headB
		}
		// 如果 p2 已经移动到了链表末尾，则将其重置为 headA，从而将其移动到第一个链表的头节点
		if p2 == nil {
			p2 = headA
		}
	}
	// 如果循环结束仍未找到交点，则说明两个链表没有交点，直接返回 nil
	return nil
}

func getIntersectionNode2(headA, headB *listnode.ListNode) *listnode.ListNode {
	// 初始化两个指针，分别指向两个链表的头节点
	p1 := headA
	p2 := headB

	// 如果两个指针都为 nil，则说明两个链表没有交点，直接返回 nil
	// 否则，执行下面的循环
	for p1 != p2 {
		// 否则，将两个指针分别向前移动一步
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
