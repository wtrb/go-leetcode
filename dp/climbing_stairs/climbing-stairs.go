package climbing

func climbStairs(n int) int {
	// return bruteForce(0, n)

	// memo := make([]int, n)
	// return memoize(memo, 0, n)

	return dypr(n)
}

// Brute force
// Time complexity: O(2^N)
// Space complexity: O(1)
func bruteForce(step int, n int) int {
	if step > n {
		return 0
	}
	if step == n {
		return 1
	}

	return bruteForce(step+1, n) + bruteForce(step+2, n)
}

// Memoize
// Time complexity: O(N)
// Space complexity: O(N)
func memoize(memo []int, step int, n int) int {
	if step > n {
		return 0
	}
	if step == n {
		return 1
	}

	if memo[step] > 0 {
		return memo[step]
	}

	memo[step] = memoize(memo, step+1, n) + memoize(memo, step+2, n)
	return memo[step]
}

// Dynamic programming
// Time complexity: O(N)
// Space complexity: O(N)
func dypr(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/

// tags: dp, dynamic programming
