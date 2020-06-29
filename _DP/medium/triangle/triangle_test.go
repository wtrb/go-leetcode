package triangle

import "testing"

func Test_minimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				triangle: [][]int{
					{2},
					{3, 4},
					{6, 5, 7},
					{4, 1, 8, 3},
				},
			},
			want: 11,
		},
		{
			name: "",
			args: args{
				triangle: [][]int{
					{2},
					{5, 1},
					{1, 7, 9},
					{1, 5, 7, 8},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
