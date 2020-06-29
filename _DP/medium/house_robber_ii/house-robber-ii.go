package robber

// Time complexity: O(N)
// Space complexity: O(1)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	return maxInts(
		rob1(nums[:len(nums)-1]),
		rob1(nums[1:]),
	)
}

// Time complexity: O(N)
// Space complexity: O(1)
func rob1(nums []int) int {
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

// House Robber II
// https://leetcode.com/problems/house-robber-ii/

// tags: dp, dynamic programming
