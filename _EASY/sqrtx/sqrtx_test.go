package sqrt

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{x: 4},
			want: 2,
		},
		{
			name: "",
			args: args{x: 8},
			want: 2,
		},
		{
			name: "",
			args: args{x: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt(%v) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}
