package twosum

import "testing"

func TestTwoSum_Find(t *testing.T) {
	ts := New()
	ts.Add(10)
	ts.Add(2)
	ts.Add(19)
	ts.Add(7)
	ts.Add(20)

	type args struct {
		twoSum *TwoSum
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				twoSum: ts,
				target: 39,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				twoSum: ts,
				target: 9,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				twoSum: ts,
				target: 10,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				twoSum: ts,
				target: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.twoSum.Find(tt.args.target); got != tt.want {
				t.Errorf("TwoSum.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
