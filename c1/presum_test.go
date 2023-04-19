package c1

import (
	"testing"
)

type NumArray struct {
	PreSums []int
	Nums    []int
}

func Constructor(nums []int) NumArray {
	// todo check

	// 计算前N个数之和，包括第N个数
	// sum[n] = sum[n-1] + num[n] 其中 sum 表示数组的前N个数之和，num此这个数组
	na := NumArray{}
	na.PreSums = make([]int, len(nums))
	na.Nums = make([]int, len(nums))
	var counter int
	for i, n := range nums {
		counter += n
		na.PreSums[i] = counter
		na.Nums[i] = n
	}
	return na
}

// 调用时需要保证 0 <= left <= right<= len(this.Nums)
func (this *NumArray) SumRange(left int, right int) int {
	if left > right {
		return -1
	}
	if left == right {
		return this.Nums[left]
	}
	return this.PreSums[right] - this.PreSums[left] + this.Nums[left]
}

// leetcode 测试用例
// ["NumArray","sumRange","sumRange","sumRange"]
// [[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]

func TestNumArray_SumRange(t *testing.T) {
	type fields struct {
		Nums []int
	}
	type args struct {
		left  []int
		right []int
	}
	tests := []struct {
		name   string //  测试用例名称
		fields fields //
		args   args   // 输入
		want   []int
	}{
		{
			name: "1",
			fields: fields{
				Nums: []int{1, 2, 3, 4, 5, 6},
			},
			args: args{
				left:  []int{0, 0, 1, 1, 1, 1, 4, 5},
				right: []int{0, 1, 1, 2, 3, 5, 5, 5},
			},
			want: []int{1, 3, 2, 5, 9, 20, 11, 6},
		},

		{
			name: "2",
			fields: fields{
				Nums: []int{-10},
			},
			args: args{
				left:  []int{0},
				right: []int{0},
			},
			want: []int{-10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.fields.Nums)
			for i := range tt.want {
				if got := this.SumRange(tt.args.left[i], tt.args.right[i]); got != tt.want[i] {
					t.Errorf("SumRange() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}
