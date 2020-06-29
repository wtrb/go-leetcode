package square

const MaxInt64 = 1<<63 - 1

func numSquares(n int) int {
	// return bruteForce(n, 1)

	// l := int(math.Sqrt(float64(n)))
	// memo := make([][]int, n+1)
	// for i := 0; i < n+1; i++ {
	// 	memo[i] = make([]int, l+1)
	// }
	// return memoize(n, 1, memo)

	return dypr(n)
}

func bruteForce(n int, i int) int {
	if n == 0 {
		return 0
	}
	if i*i > n {
		return MaxInt64
	}

	exclusive := bruteForce(n, i+1)
	inclusive := 1 + bruteForce(n-i*i, 1)

	return minInts(inclusive, exclusive)
}

func memoize(n int, i int, memo [][]int) int {
	if n == 0 {
		return 0
	}
	if i*i > n {
		return MaxInt64
	}

	if memo[n][i] != 0 {
		return memo[n][i]
	}

	exclusive := memoize(n, i+1, memo)
	inclusive := 1 + memoize(n-i*i, 1, memo)

	memo[n][i] = minInts(inclusive, exclusive)

	return memo[n][i]
}

func dypr(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = MaxInt64
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = minInts(
				dp[i],
				1+dp[i-j*j],
			)
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

// https://stackoverflow.com/questions/39031099/perfect-square-in-leetcode

// tags: dp, dynamic programming
