package move_zeroes_283

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				[]int{},
			},
			want: []int{},
		},
		{
			name: "",
			args: args{
				[]int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "",
			args: args{
				nums: []int{12, 0, 0},
			},
			want: []int{12, 0, 0},
		},
		{
			name: "",
			args: args{
				nums: []int{0, 1, 0},
			},
			want: []int{1, 0, 0},
		},
		{
			name: "",
			args: args{
				nums: []int{0, 0, 2},
			},
			want: []int{2, 0, 0},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 0, 2},
			},
			want: []int{1, 2, 0},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: []int{1, 1, 1},
		},
		{
			name: "",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: []int{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			got := tt.args.nums

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
