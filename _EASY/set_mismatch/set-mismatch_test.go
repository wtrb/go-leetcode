package set

import (
	"reflect"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
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
			args: args{nums: []int{1, 2, 2, 4}},
			want: []int{2, 3},
		},
		{
			name: "",
			args: args{nums: []int{2, 3}},
			want: []int{2, 1},
		},
		{
			name: "",
			args: args{nums: []int{1, 3, 3}},
			want: []int{3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findErrorNums(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums(%+v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
