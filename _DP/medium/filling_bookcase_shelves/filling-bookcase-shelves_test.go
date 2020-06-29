package bookcase

import "testing"

func Test_minHeightShelves(t *testing.T) {
	type args struct {
		books      [][]int
		shelfWidth int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				books:      [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}},
				shelfWidth: 4,
			},
			want: 6,
		},
		// {
		// 	name: "",
		// 	args: args{
		// 		books:      [][]int{},
		// 		shelfWidth: 0,
		// 	},
		// 	want: 0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minHeightShelves(tt.args.books, tt.args.shelfWidth); got != tt.want {
				t.Errorf("minHeightShelves() = %v, want %v", got, tt.want)
			}
		})
	}
}
