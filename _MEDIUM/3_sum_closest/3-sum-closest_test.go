package threesum

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 1, 1, 0},
				target: -100,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 1, -1, -1, 3},
				target: 1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums:   []int{-1, 0, 1, 2, -1, -4},
				target: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest(%+v, %v) = %v, want %v", tt.args.nums, tt.args.target, got, tt.want)
			}
		})
	}
}
