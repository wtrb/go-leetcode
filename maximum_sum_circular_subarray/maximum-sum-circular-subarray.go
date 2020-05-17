package cirarray

// Time complexity: O(N)
// Space complexity: O(1)
func maxSubarraySumCircular(A []int) int {
	// #case 1
	maxSum1 := A[0]
	curSum1 := A[0]
	for i := 1; i < len(A); i++ {
		curSum1 = maxInts(curSum1, 0) + A[i]
		maxSum1 = maxInts(maxSum1, curSum1)
	}

	// #case 2
	maxLeftArr := make([]int, len(A))
	maxLeftArr[len(A)-1] = A[len(A)-1]
	leftSum := A[len(A)-1]
	for i := len(A) - 2; i >= 1; i-- {
		leftSum += A[i]
		maxLeftArr[i] = maxInts(maxLeftArr[i+1], leftSum)
	}

	rightSumArr := make([]int, len(A))
	rightSumArr[0] = A[0]
	for i := 1; i < len(A)-1; i++ {
		rightSumArr[i] = rightSumArr[i-1] + A[i]
	}

	maxSum2 := A[0]
	for i := 0; i < len(A)-1; i++ {
		curMax2 := rightSumArr[i] + maxLeftArr[i+1]
		maxSum2 = maxInts(maxSum2, curMax2)
	}

	// result
	return maxInts(maxSum1, maxSum2)
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Maximum Sum Circular Subarray
// https://leetcode.com/problems/maximum-sum-circular-subarray
// https://leetcode.com/submissions/detail/340477972/?from=/explore/featured/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3330/

// https://www.youtube.com/watch?v=vxer7ns7BIs

// tags: kadane, max slice, max sum subarray, dynamic programming
