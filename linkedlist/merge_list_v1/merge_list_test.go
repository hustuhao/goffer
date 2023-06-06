package merge_list_v1

import "testing"

func TestMergeKLists(t *testing.T) {
	// 用于构建链表
	newList := func(vals []int) *ListNode {
		dummy := &ListNode{}
		cur := dummy
		for _, val := range vals {
			cur.Next = &ListNode{Val: val}
			cur = cur.Next
		}
		return dummy.Next
	}

	// 测试用例
	testCases := []struct {
		input []*ListNode
		want  []int
	}{
		{
			input: []*ListNode{newList([]int{1, 4, 5}), newList([]int{1, 3, 4}), newList([]int{2, 6})},
			want:  []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			input: []*ListNode{newList([]int{1, 2, 3}), newList([]int{}), newList([]int{4, 5, 6})},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			input: []*ListNode{newList([]int{1, 2, 3}), newList([]int{4, 5, 6}), newList([]int{7, 8, 9})},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},

		{
			input: []*ListNode{newList([]int{1}), newList([]int{0})},
			want:  []int{0, 1},
		},
	}

	// 循环进行测试
	for i, tc := range testCases {
		//got := mergeKLists(tc.input)
		got := mergeKListsV2(tc.input)
		if !compareLinkedList(got, newList(tc.want)) {
			t.Errorf("mergeKLists(%v) = %v, want %v", tc.input, toSlice(got), tc.want)
		}
		t.Logf("case %d passed", i+1)
	}
}

// 判断两个链表是否相等
func compareLinkedList(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// 将链表转化为数组
func toSlice(l *ListNode) []int {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}
