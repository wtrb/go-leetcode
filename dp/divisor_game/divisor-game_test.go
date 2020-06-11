package dvgame

import "testing"

func Test_divisorGame(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				N: 2,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				N: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorGame(tt.args.N); got != tt.want {
				t.Errorf("divisorGame(%v) = %v, want %v", tt.args.N, got, tt.want)
			}
		})
	}
}
