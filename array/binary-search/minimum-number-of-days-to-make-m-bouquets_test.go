package binary_search

import (
	"testing"
)

func TestMinDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				bloomDay: []int{1, 2, 3, 4, 5},
				m:        3,
				k:        1,
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				bloomDay: []int{1, 1, 1, 1, 1},
				m:        3,
				k:        1,
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				bloomDay: []int{1, 1, 1, 1, 1},
				m:        2,
				k:        2,
			},
			want: 1,
		},
		{
			name: "case4",
			args: args{
				bloomDay: []int{1, 1, 1, 1, 1},
				m:        2,
				k:        3,
			},
			want: -1,
		},
		{
			name: "case5",
			args: args{
				bloomDay: []int{1, 1, 1, 1, 1},
				m:        2,
				k:        2,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args.bloomDay, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
