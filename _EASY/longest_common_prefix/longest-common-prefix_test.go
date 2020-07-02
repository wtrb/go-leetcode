package string

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "",
			args: args{
				strs: []string{},
			},
			want: "",
		},
		{
			name: "",
			args: args{
				strs: []string{"abc"},
			},
			want: "abc",
		},
		{
			name: "",
			args: args{
				strs: []string{"abc", "ade"},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix(%+v) = %v, want %v", tt.args.strs, got, tt.want)
			}
		})
	}
}
