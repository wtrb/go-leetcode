package roman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{num: 5},
			want: "V",
		},
		{
			name: "",
			args: args{num: 10},
			want: "X",
		},
		{
			name: "",
			args: args{num: 3},
			want: "III",
		},
		{
			name: "",
			args: args{num: 4},
			want: "IV",
		},
		{
			name: "",
			args: args{num: 9},
			want: "IX",
		},
		{
			name: "",
			args: args{num: 58},
			want: "LVIII",
		},
		{
			name: "",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman(%v) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}
