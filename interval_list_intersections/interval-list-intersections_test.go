package list

import (
	"reflect"
	"testing"
)

func Test_intervalIntersection(t *testing.T) {
	type args struct {
		A [][]int
		B [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				A: [][]int{{1, 3}},
				B: [][]int{{2, 4}},
			},
			want: [][]int{{2, 3}},
		},
		{
			name: "",
			args: args{
				A: [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
				B: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			},
			want: [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalIntersection(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
