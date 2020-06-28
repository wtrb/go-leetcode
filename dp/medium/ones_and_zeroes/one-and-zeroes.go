package oz

func findMaxForm(strs []string, m int, n int) int {
	// return recursion(strs, 0, m, n)

	// memo := make([][][]int, len(strs))
	// for i := 0; i < len(strs); i++ {
	// 	memo[i] = make([][]int, m+1)
	// 	for k := 0; k < m+1; k++ {
	// 		memo[i][k] = make([]int, n+1)
	// 	}
	// }
	// return memoize(strs, 0, m, n, memo)

	// return dypr3D(strs, m, n)

	return dypr2D(strs, m, n)
}

// Time Complexity: O(2^len * s) - len is Strs length, s is the average string length
// Space Complexity: O(1)
func recursion(strs []string, i int, m, n int) int {
	if i == len(strs) {
		return 0
	}

	zeroes, ones := countZeroOnes(strs[i])

	in := 0
	if m >= zeroes && n >= ones {
		in = 1 + recursion(strs, i+1, m-zeroes, n-ones)
	}

	ex := recursion(strs, i+1, m, n)

	return maxInts(ex, in)
}

// Time Complexity: O(l * m * n) - l is length of strs, m is number of 0s, n is number of 1s
// Space Complexity: O(l * m * n) - memo 3D Array
func memoize(strs []string, i int, m, n int, memo [][][]int) int {
	if i == len(strs) {
		return 0
	}

	if memo[i][m][n] != 0 {
		return memo[i][m][n]
	}

	zeroes, ones := countZeroOnes(strs[i])

	if m >= zeroes && n >= ones {
		memo[i][m][n] = 1 + memoize(strs, i+1, m-zeroes, n-ones, memo)
	}

	memo[i][m][n] = maxInts(
		memo[i][m][n],
		memoize(strs, i+1, m, n, memo),
	)

	return memo[i][m][n]
}

func dypr3D(strs []string, m, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := 0; i < len(strs)+1; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, n+1)
			dp[i][j][0] = 0
		}
	}

	for i := 1; i < len(strs)+1; i++ {
		zeros, ones := countZeroOnes(strs[i-1])

		for j := 0; j < m+1; j++ {
			for k := 0; k < n+1; k++ {
				if j >= zeros && k >= ones {
					dp[i][j][k] = maxInts(dp[i-1][j][k], 1+dp[i-1][j-zeros][k-ones])

				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}

func dypr2D(strs []string, m, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		zeroes, ones := countZeroOnes(strs[i])

		for j := m; j >= zeroes; j-- {
			for k := n; k >= ones; k-- {
				dp[j][k] = maxInts(
					dp[j][k],
					1+dp[j-zeroes][k-ones],
				)
			}
		}
	}

	return dp[m][n]
}

func countZeroOnes(str string) (zeros, ones int) {
	for _, c := range str {
		if c == '0' {
			zeros++

		} else if c == '1' {
			ones++
		}
	}

	return
}

func maxInts(ints ...int) int {
	max := ints[0]
	for _, v := range ints {
		if v > max {
			max = v
		}
	}
	return max
}

// Ones and Zeroes
// https://leetcode.com/problems/ones-and-zeroes/

// https://github.com/azl397985856/leetcode/blob/master/problems/474.ones-and-zeros-en.md

// tags: dy, dynamic programming, Knapsack
