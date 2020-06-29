package coin

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				coins:  []int{2},
				amount: 3,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				coins:  []int{346, 29, 395, 188, 155, 109},
				amount: 9401,
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange(%+v, %v) = %v, want %v", tt.args.coins, tt.args.amount, got, tt.want)
			}
		})
	}
}
