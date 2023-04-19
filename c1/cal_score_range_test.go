package c1

import (
	"testing"
)

// 分析一个班级学生的考试成绩（0-100分），给出成绩范围，查询在此区间段内的人数
// 关键词：优先范围，人数和
func TestAnalysisScoreRangeNum(t *testing.T) {
	type fields struct {
		Scores []int
	}
	type args struct {
		Left  int
		Right int
		Want  int
	}

	tests := []struct {
		name   string
		fields fields
		args   []args
	}{
		{
			name: "用例1",
			fields: fields{
				Scores: []int{1, 2, 40, 12, 48, 58, 88, 12, 3, 88, 99, 100},
			},
			args: []args{
				{
					Left:  0,
					Right: 0,
					Want:  0,
				},
				{
					Left:  1,
					Right: 1,
					Want:  1,
				},
				{
					Left:  1,
					Right: 2,
					Want:  2,
				},
				{
					Left:  0,
					Right: 1,
					Want:  1,
				},
				{
					Left:  0,
					Right: 100,
					Want:  12,
				},
				{
					Left:  100,
					Right: 100,
					Want:  1,
				},
				{
					Left:  50,
					Right: 100,
					Want:  5,
				},
			},
		},
	}

	for _, tt := range tests {
		sa := ScoreArrayConstructor(tt.fields.Scores)
		for _, arg := range tt.args {
			if got := sa.SumRange(arg.Left, arg.Right); got != arg.Want {
				t.Fatalf("case:%s, arg:%v, want:%v, got:%v", tt.name, arg, arg.Want, got)
			}
		}
	}
}

// //先计算每个分数有多少个同学，再利用前缀和和分数来计算分数段区间内的人数
type ScoreArray struct {
	ScoresPreSum []int
}

func ScoreArrayConstructor(scores []int) ScoreArray {
	dst := make([]int, 101) // 100分的满分
	for i := 0; i < len(scores); i++ {
		dst[scores[i]]++
	}

	for i := 1; i < len(dst); i++ {
		dst[i] = dst[i-1] + dst[i]
	}
	return ScoreArray{
		ScoresPreSum: dst,
	}
}

// left 和 right 必须在 [0,100]
func (this *ScoreArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.ScoresPreSum[right]
	}
	return this.ScoresPreSum[right] - this.ScoresPreSum[left-1]
}
