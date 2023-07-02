package binary_search

import "testing"

func TestSplitArrayLargest(t *testing.T) {
	type arg struct {
		nums []int
		k    int
	}

	tests := []struct {
		Name string
		Arg  arg
		Want int
	}{
		{
			Name: "case 1",
			Arg: arg{
				nums: []int{7, 2, 5, 10, 8},
				k:    2,
			},
			Want: 18,
		}, {
			Name: "case 2",
			Arg: arg{
				nums: []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
				k:    1,
			},
			Want: 54,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := splitArray(tt.Arg.nums, tt.Arg.k); got != tt.Want {
				t.Fatalf("expect:%d, got:%d", tt.Want, got)
			}
		})

	}

}
