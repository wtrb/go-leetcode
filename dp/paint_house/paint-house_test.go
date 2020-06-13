package paint

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				costs: [][]int{
					{17, 2, 17},
					{16, 16, 5},
					{14, 3, 19},
				},
			},
			want: 10,
		},
		{
			name: "",
			args: args{
				costs: [][]int{
					{1, 2, 3},
					{7, 4, 5},
					{6, 9, 2},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.costs); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
