package square

func numSquares(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = 1<<63 - 1
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = minInts(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Perfect Squares
// https://leetcode.com/problems/perfect-squares/

// tags: dp, dynamic programming
