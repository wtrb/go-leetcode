package paint

import "testing"

func Test_paintFence(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n: 0,
				k: 0,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				n: 1,
				k: 5,
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				n: 3,
				k: 2,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				n: 2,
				k: 4,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paintFence(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("paintFence(%v, %v) = %v, want %v", tt.args.n, tt.args.k, got, tt.want)
			}
		})
	}
}
