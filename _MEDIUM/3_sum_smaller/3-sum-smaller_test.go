package threesum

import "testing"

func Test_threeSumSmaller(t *testing.T) {
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
				nums:   []int{-2, 0, 1, 3},
				target: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumSmaller(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumSmaller(%+v, %v) = %v, want %v", tt.args.nums, tt.args.target, got, tt.want)
			}
		})
	}
}
