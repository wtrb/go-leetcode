package product

import "testing"

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 24,
		},
		{
			name: "",
			args: args{
				nums: []int{-2, 1, 2, 3, 4},
			},
			want: 24,
		},
		{
			name: "",
			args: args{
				nums: []int{-5, -2, 1, 2, 3, 4},
			},
			want: 40,
		},
		{
			name: "",
			args: args{
				nums: []int{-3, 1, 2, -2, 5, 6},
			},
			want: 60,
		},
		{
			name: "",
			args: args{
				nums: []int{-5, 5, -5, 4},
			},
			want: 125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct(%+v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
