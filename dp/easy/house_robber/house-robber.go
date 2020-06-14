package robber

func rob(nums []int) int {
	// return maxInts(bruteForce(nums, 0, true), bruteForce(nums, 0, false))
	return dypr(nums)
}

// Brute force
// Time complexity: O(2^N)
// Space complexity: O(1)
func bruteForce(nums []int, cur int, skip bool) int {
	if cur > len(nums)-1 {
		return 0
	}

	if skip {
		return maxInts(bruteForce(nums, cur+1, false), bruteForce(nums, cur+1, true))
	}

	return nums[cur] + bruteForce(nums, cur+1, true)
}

// Dynamic programming
// Time complexity: O(N)
// Space complexity: O(N)
func dypr(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// if len(nums) == 1 {
	// 	return nums[0]
	// }

	// dp := make([]int, len(nums))
	// dp[0] = nums[0]                   // rob 1st house -> 1st house's money
	// dp[1] = maxInts(nums[0], nums[1]) // rob 2nd house -> max(2nd, 1st)
	// for i := 2; i < len(nums); i++ {
	// 	dp[i] = maxInts(nums[i]+dp[i-2], dp[i-1]) // rob 3rd house -> max(3rd+1st, 2nd)
	// }
	// return dp[len(nums)-1]

	dp := make([]int, len(nums)+1)
	dp[0] = 0       // rob 0 house -> 0 money
	dp[1] = nums[0] // rob 1st house -> 1st house's money
	for i := 1; i < len(nums); i++ {
		dp[i+1] = maxInts(dp[i], dp[i-1]+nums[i])
	}
	return dp[len(nums)]
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// House Robber
// https://leetcode.com/problems/house-robber/

// https://www.youtube.com/watch?v=xlvhyfcoQa4

// tags: dp, dynamic programming
