package coin

import "testing"

func Test_arrangeCoins(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n: 5,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				n: 8,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangeCoins(tt.args.n); got != tt.want {
				t.Errorf("arrangeCoins(%+v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
