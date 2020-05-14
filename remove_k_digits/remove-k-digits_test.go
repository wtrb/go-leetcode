package digits

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				num: "1432219",
				k:   3,
			},
			want: "1219",
		},
		{
			name: "",
			args: args{
				num: "10200",
				k:   1,
			},
			want: "200",
		},
		{
			name: "",
			args: args{
				num: "10",
				k:   2,
			},
			want: "0",
		},
		{
			name: "",
			args: args{
				num: "100200",
				k:   2,
			},
			want: "0",
		},
		{
			name: "",
			args: args{
				num: "1002000",
				k:   2,
			},
			want: "0",
		},
		{
			name: "",
			args: args{
				num: "10",
				k:   1,
			},
			want: "0",
		},
		{
			name: "",
			args: args{
				num: "112",
				k:   1,
			},
			want: "11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits(\"%v\", %v) = %v, want %v", tt.args.num, tt.args.k, got, tt.want)
			}
		})
	}
}
