package maxsub

// Time complexity: O(n)
// Space complexity: O(1)
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum = maxInts(sum+nums[i], nums[i])
		maxSum = maxInts(maxSum, sum)
	}

	return maxSum
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Maximum Subarray
// https://leetcode.com/problems/maximum-subarray/

// tags: max sum, dp
