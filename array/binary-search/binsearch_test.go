package binary_search

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		list   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1: empty list",
			args: args{
				list:   nil,
				target: 0,
			},
			want: -1,
		},
		{
			name: "case 2",
			args: args{
				list:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				list:   []int{1},
				target: 2,
			},
			want: -1,
		},
		{
			name: "case 4",
			args: args{
				list:   []int{1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "case 5",
			args: args{
				list:   []int{2, 4, 6, 8, 10},
				target: 3,
			},
			want: -1,
		},
		{
			name: "case 6",
			args: args{
				list:   []int{2, 4, 6, 8, 10},
				target: 8,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leftBound(t *testing.T) {
	type args struct {
		list   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				list:   []int{1, 2, 3},
				target: 2,
			},
			want: 1,
		},

		{
			name: "case 2",
			args: args{
				list:   []int{1, 2, 2, 3},
				target: 2,
			},
			want: 1,
		},

		{
			name: "case 3",
			args: args{
				list:   []int{1, 2, 2, 3},
				target: 3,
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				list:   []int{1, 2, 2, 3},
				target: 4,
			},
			want: -1,
		},
		{
			name: "case 5",
			args: args{
				list:   []int{1, 2, 2, 3},
				target: 0,
			},
			want: -1,
		},
		{
			name: "case 6",
			args: args{
				list:   []int{1, 2, 2, 2, 3},
				target: 2,
			},
			want: 1,
		},
		{
			name: "case 7",
			args: args{
				list:   []int{1, 2, 2, 2, 3, 3},
				target: 3,
			},
			want: 4,
		},
		{
			name: "case 8",
			args: args{
				list:   []int{1, 1, 2, 2, 2, 3, 3},
				target: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftBound(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("leftBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rightBound(t *testing.T) {
	type args struct {
		list   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				list:   []int{1, 2, 3},
				target: 2,
			},
			want: 1,
		},

		{
			name: "case 2",
			args: args{
				list:   []int{1, 2, 2, 3},
				target: 2,
			},
			want: 2,
		},

		{
			name: "case 3",
			args: args{
				list:   []int{1, 2, 2, 3},
				target: 3,
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				list:   []int{1, 2, 2, 3},
				target: 4,
			},
			want: -1,
		},
		{
			name: "case 5",
			args: args{
				list:   []int{1, 2, 2, 3},
				target: 0,
			},
			want: -1,
		},
		{
			name: "case 6",
			args: args{
				list:   []int{1, 2, 2, 2, 3},
				target: 2,
			},
			want: 3,
		},
		{
			name: "case 7",
			args: args{
				list:   []int{1, 2, 2, 2, 3, 3},
				target: 3,
			},
			want: 5,
		},
		{
			name: "case 8",
			args: args{
				list:   []int{1, 1, 2, 2, 2, 3, 3},
				target: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightBound(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("rightBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
