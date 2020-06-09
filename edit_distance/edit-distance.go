package distance

// Time complexity: O(m*n)
// Space complexity: O(m*n)
func minDistance(source string, target string) int {
	m, n := len(source), len(target)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// If we only to DELETE source
	for i := 1; i < m+1; i++ {
		dp[i][0] = i
	}
	// then, INSERT (copy target),
	for j := 1; j < n+1; j++ {
		dp[0][j] = j
	}
	// we would be done here, but it's not minimum distance.
	// We need to combine with REPLACE/SUBSTITUTION.
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if source[i-1] == target[j-1] {
				dp[i][j] = dp[i-1][j-1]

			} else {
				dp[i][j] = 1 + // need 1 replace/substitution operation here
					minInts(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}

func minInts(nums ...int) int {
	min := 1<<63 - 1
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

// Edit Distance
// https://leetcode.com/problems/edit-distance/

// https://www.youtube.com/watch?v=gmf5RY8s3QE
// https://www.hackerearth.com/practice/algorithms/dynamic-programming/2-dimensional/tutorial/

// tags: dp, dynamic programming
