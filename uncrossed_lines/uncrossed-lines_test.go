package lines

import "testing"

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				A: []int{1, 4, 2},
				B: []int{1, 2, 4},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				A: []int{2, 5, 1, 2, 5},
				B: []int{10, 5, 2, 1, 5, 2},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				A: []int{1, 3, 7, 1, 7, 5},
				B: []int{1, 9, 2, 5, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("maxUncrossedLines(A: %+v, B: %+v) = %v, want %v", tt.args.A, tt.args.B, got, tt.want)
			}
		})
	}
}
