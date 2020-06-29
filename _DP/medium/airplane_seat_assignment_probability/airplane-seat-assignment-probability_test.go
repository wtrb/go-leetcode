package probability

import "testing"

func Test_nthPersonGetsNthSeat(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{
				n: 1,
			},
			want: 1.0,
		},
		{
			name: "",
			args: args{
				n: 2,
			},
			want: 0.5,
		},
		{
			name: "",
			args: args{
				n: 3,
			},
			want: 0.5,
		},
		{
			name: "",
			args: args{
				n: 4,
			},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthPersonGetsNthSeat(tt.args.n); got != tt.want {
				t.Errorf("nthPersonGetsNthSeat(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
