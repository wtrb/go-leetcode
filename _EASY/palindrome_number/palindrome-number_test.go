package number

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				x: 10,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				x: 0,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				x: 120,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome(%v) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}
