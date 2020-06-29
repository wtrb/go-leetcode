package predict

func PredictTheWinner(nums []int) bool {
	// return recursion(nums, 0, len(nums)-1) >= 0

	// memo := make([][]int, len(nums))
	// for i := 0; i < len(nums); i++ {
	// 	memo[i] = make([]int, len(nums))
	// 	for j := 0; j < len(nums); j++ {
	// 		memo[i][j] = -1<<63
	// 	}
	// }
	// return memoize(nums, 0, len(nums)-1, memo) >= 0

	return dypr(nums)
}

func recursion(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	return maxInts(
		nums[left]-recursion(nums, left+1, right),
		nums[right]-recursion(nums, left, right-1),
	)
}

func memoize(nums []int, left, right int, memo [][]int) int {
	if left == right {
		return nums[left]
	}

	if memo[left][right] != -1<<63 {
		return memo[left][right]
	}

	memo[left][right] = maxInts(
		nums[left]-memoize(nums, left+1, right, memo),
		nums[right]-memoize(nums, left, right-1, memo),
	)

	return memo[left][right]
}

// Time complexity: O(n^2)
// Space complexity: O(n^2)
func dypr(nums []int) bool {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			dp[i][j] = maxInts(
				nums[i]-dp[i+1][j],
				nums[j]-dp[i][j-1],
			)
		}
	}

	return dp[0][len(nums)-1] >= 0
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Predict the Winner
// https://leetcode.com/problems/predict-the-winner/
// Similar to https://leetcode.com/problems/stone-game/

// tags: dp, dynamic programming
