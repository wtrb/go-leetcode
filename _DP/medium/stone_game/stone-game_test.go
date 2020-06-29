package game

import "testing"

func Test_stoneGame(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				piles: []int{5, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				piles: []int{1, 5, 233, 7},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGame(tt.args.piles); got != tt.want {
				t.Errorf("stoneGame(%+v) = %v, want %v", tt.args.piles, got, tt.want)
			}
		})
	}
}
