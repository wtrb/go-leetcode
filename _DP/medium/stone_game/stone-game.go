package game

func stoneGame(piles []int) bool {
	// return loop(piles)
	// return math(piles)
	// return recursion(piles, 0, len(piles)-1) > 0

	// memo := make([][]int, len(piles))
	// for i := 0; i < len(piles); i++ {
	// 	memo[i] = make([]int, len(piles))
	// 	for j := 0; j < len(piles); j++ {
	// 		memo[i][j] = -1 << 63
	// 	}
	// }
	// return memoize(piles, 0, len(piles)-1, memo) > 0

	return dypr(piles)
}

// Time complexity: O(n)
// Space complexity: O(1)
func loop(piles []int) bool {
	alex := 0
	lee := 0
	for len(piles) > 0 {
		head := piles[0]
		tail := piles[len(piles)-1]

		if head > tail {
			alex += head
			lee += tail
		} else {
			alex += tail
			lee += head
		}
		piles = piles[1 : len(piles)-1]
	}

	if alex > lee {
		return true
	}
	return false
}

// Time complexity: O(1)
// Space complexity: O(1)
// Alex starts first and can always pick the pile with more stones.
// Therefore, Alex will always win this game.
// LOL.
func math(piles []int) bool {
	return true
}

// Time complexity: O(2^n)
// Space complexity: O(n)
func recursion(piles []int, left, right int) int {
	if left == right {
		return piles[left]
	}

	return maxInts(piles[left]-recursion(piles, left+1, right), piles[right]-recursion(piles, left, right-1))
}

// Time complexity: O(n^2)
// Space complexity: O(n^2)
func memoize(piles []int, left, right int, memo [][]int) int {
	if left == right {
		return piles[left]
	}

	if memo[left][right] == -1<<63 {
		memo[left][right] = maxInts(
			piles[left]-memoize(piles, left+1, right, memo),
			piles[right]-memoize(piles, left, right-1, memo),
		)
	}

	return memo[left][right]
}

// Time complexity: O(n^2)
// Space complexity: O(n^2)
func dypr(piles []int) bool {
	dp := make([][]int, len(piles))
	for i := 0; i < len(piles); i++ {
		dp[i] = make([]int, len(piles))
		dp[i][i] = piles[i]
	}

	for i := 0; i < len(piles); i++ {
		for j := i + 1; j < len(piles); j++ {
			dp[i][j] = maxInts(
				piles[i]-dp[i+1][j],
				piles[j]-dp[i][j-1],
			)
		}
	}

	return dp[0][len(piles)-1] > 0
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Stone Game
// https://leetcode.com/problems/stone-game/
// Similar to https://leetcode.com/problems/predict-the-winner/

// tags: dp, dynamic programming, math, recursion, memoize
