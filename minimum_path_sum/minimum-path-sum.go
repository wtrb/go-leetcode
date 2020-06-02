package sum

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	dp := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))

		for j := 0; j < len(grid[0]); j++ {
			dp[i][j] += grid[i][j]

			if i > 0 && j > 0 {
				dp[i][j] += func(a, b int) int {
					if a < b {
						return a
					}
					return b
				}(dp[i][j-1], dp[i-1][j])

			} else if i > 0 {
				dp[i][j] += dp[i-1][j]

			} else if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

// Minimum Path Sum
// https://leetcode.com/problems/minimum-path-sum/
// https://leetcode.com/submissions/detail/326924947/?from=/explore/featured/card/30-day-leetcoding-challenge/530/week-3/3303/

// tags: dp, dynamic programming
