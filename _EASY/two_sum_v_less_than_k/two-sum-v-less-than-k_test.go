package twosum

import "testing"

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				numbers: []int{34, 23, 1, 24, 75, 33, 54, 8},
				k:       60,
			},
			want: 58,
		},
		{
			name: "",
			args: args{
				numbers: []int{10, 20, 30},
				k:       15,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.k); got != tt.want {
				t.Errorf("twoSum(%+v, %v) = %v, want %v", tt.args.numbers, tt.args.k, got, tt.want)
			}
		})
	}
}
