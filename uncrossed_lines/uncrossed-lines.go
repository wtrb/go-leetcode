package lines

import "fmt"

func maxUncrossedLines(A []int, B []int) int {
	m, n := len(A), len(B)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				
			} else {
				dp[i][j] = maxInts(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	fmt.Printf("%+v", dp)

	return dp[m][n]
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Uncrossed Lines
// https://leetcode.com/problems/uncrossed-lines/

// https://leetcode.com/submissions/detail/344789739/?from=/explore/featured/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3340/

// tags: dp, dynamic programming
