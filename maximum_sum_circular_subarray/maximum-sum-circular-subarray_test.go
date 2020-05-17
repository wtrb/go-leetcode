package cirarray

import "testing"

func Test_maxSubarraySumCircular(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				A: []int{1, -2, 3, -2},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				A: []int{5, -3, 5},
			},
			want: 10,
		},
		{
			name: "",
			args: args{
				A: []int{3, -1, 2, -1},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				A: []int{3, -2, 2, -3},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				A: []int{-2, -3, -1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircular(tt.args.A); got != tt.want {
				t.Errorf("maxSubarraySumCircular(%v) = %v, want %v", tt.args.A, got, tt.want)
			}
		})
	}
}
