package revert

import (
	"goffer/common/listnode"
	"testing"
)

func TestRevertList(t *testing.T) {
	testCases := []struct {
		linkedList []int
		want       int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			5,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{1},
			1,
		},
		{
			nil,
			-1,
		},
	}

	for i := 0; i < len(testCases); i++ {
		head := listnode.NewList(testCases[i].linkedList)
		newHead := reverseList3(head)
		if newHead == nil {
			if testCases[i].want != -1 {
				t.Fail()
			}
		} else if newHead.Val != testCases[i].want {
			t.Fail()
		}
	}
}
