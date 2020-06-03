package square

// Time complexity : O(mn)
// Space complexity : O(mn)
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	rows := len(matrix)
	cols := len(matrix[0])
	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, cols+1)
	}

	count := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + minInts(dp[i-1][j-1], minInts(dp[i-1][j], dp[i][j-1]))

				count = maxInts(count, dp[i][j])
			}
		}
	}

	return count * count
}

func minInts(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInts(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Maximal Square
// https://leetcode.com/problems/maximal-square/
// https://leetcode.com/submissions/detail/331151459/?from=/explore/featured/card/30-day-leetcoding-challenge/531/week-4/3312/

// tags: dp, danymic programming
