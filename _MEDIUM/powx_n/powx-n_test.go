package pow

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{x: 2.0, n: 0},
			want: 1.00000,
		},
		{
			name: "",
			args: args{x: 0.0, n: 10},
			want: 0.00000,
		},
		{
			name: "",
			args: args{x: 2.1, n: 3},
			want: 9.261000000000001,
		},
		{
			name: "",
			args: args{x: 2.0, n: -2},
			want: 0.25000,
		},
		{
			name: "",
			args: args{x: 34.00515, n: -3},
			want: 2.543114507074558e-05,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow(%v, %v) = %v, want %v", tt.args.x, tt.args.n, got, tt.want)
			}
		})
	}
}
