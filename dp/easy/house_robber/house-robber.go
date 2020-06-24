package robber

func rob(nums []int) int {
	// return recursion(nums, 0)

	// memo := make([]int, len(nums))
	// return memoize(nums, 0, memo)

	// return dypr(nums)

	return dyprO1(nums)
}

// Time complexity: O(2^N)
// Space complexity: O(1)
func recursion(nums []int, i int) int {
	if i >= len(nums) {
		return 0
	}

	return maxInts(
		nums[i]+recursion(nums, i+2), // rob, then skip i+1-th house
		recursion(nums, i+1),         // don't rob, process to i+1-th house
	)
}

// Time complexity: O(N)
// Space complexity: O(N)
func memoize(nums []int, i int, memo []int) int {
	if i >= len(nums) {
		return 0
	}

	if memo[i] != 0 {
		return memo[i]
	}

	memo[i] = maxInts(
		nums[i]+memoize(nums, i+2, memo), // rob, then skip i+1-th house
		memoize(nums, i+1, memo),         // don't rob, process to i+1-th house
	)

	return memo[i]
}

// Time complexity: O(N)
// Space complexity: O(N)
func dypr(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i+1] = maxInts(
			dp[i],           // don't rob
			dp[i-1]+nums[i], // rob
		)
	}

	return dp[len(nums)]
}

// Time complexity: O(N)
// Space complexity: O(1)
func dyprO1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	a := 0
	b := nums[0]
	for i := 1; i < len(nums); i++ {
		a, b = b, maxInts(b, a+nums[i])
	}

	return b
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// House Robber
// https://leetcode.com/problems/house-robber/

// tags: dp, dynamic programming
