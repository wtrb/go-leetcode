package binarymatrix

type BinaryMatrix interface {
	Get(x, y int) int
	Dimensions() []int
}

// Time complexity: O(n+m)
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	di := binaryMatrix.Dimensions()
	rows := di[0]
	cols := di[1]

	row := 0
	col := cols - 1
	for row < rows && col >= 0 {
		if binaryMatrix.Get(row, col) == 0 {
			row++

		} else {
			col--
		}
	}

	if col != cols-1 {
		return col + 1

	} else {
		return -1
	}
}

// Leftmost Column with at Least a One

// https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/530/week-3/3306/
// https://leetcode.com/submissions/detail/328383330/?from=/explore/featured/card/30-day-leetcoding-challenge/530/week-3/3306/

// tags:
