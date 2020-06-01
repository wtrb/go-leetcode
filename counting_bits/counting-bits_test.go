package bits

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				num: 2,
			},
			want: []int{0, 1, 1},
		},
		{
			name: "",
			args: args{
				num: 5,
			},
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits(%v) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}
