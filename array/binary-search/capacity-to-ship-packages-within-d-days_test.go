package binary_search

import "testing"

func TestShipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		day     int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				weights: []int{1, 2, 3, 1, 1},
				day:     4,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				weights: []int{2, 2, 2, 2, 2, 2},
				day:     4,
			},
			want: 4,
		}, {
			name: "",
			args: args{
				weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				day:     5,
			},
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shipWithinDays(tt.args.weights, tt.args.day); got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})

	}

}
