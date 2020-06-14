package string

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s: "",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence(%v, %v) = %v, want %v", tt.args.s, tt.args.t, got, tt.want)
			}
		})
	}
}
