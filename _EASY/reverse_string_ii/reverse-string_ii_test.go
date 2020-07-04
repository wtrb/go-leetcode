package string

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{s: "abcdefg", k: 2},
			want: "bacdfeg",
		},
		{
			name: "",
			args: args{s: "abcdefg", k: 3},
			want: "cbadefg",
		},
		{
			name: "",
			args: args{s: "abcdefgh", k: 3},
			want: "cbadefhg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr(%v, %v) = %v, want %v", tt.args.s, tt.args.k, got, tt.want)
			}
		})
	}
}
