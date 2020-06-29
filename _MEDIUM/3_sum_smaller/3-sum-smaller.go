package threesum

import (
	"sort"
)

func threeSumSmaller(nums []int, target int) int {
	count := 0

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		lo, hi := i+1, len(nums)-1

		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]

			if sum < target {
				count += (hi - lo) // (hi-lo) pairs < target
				lo++

			} else {
				hi--
			}
		}
	}

	return count
}

// 3Sum Smaller
// https://leetcode.com/problems/3sum-smaller/

/*
	Given an array of n integers nums and a target, find the number of index triplets i, j, k with 0 <= i < j < k < n that satisfy the condition nums[i] + nums[j] + nums[k] < target.

	Example:
	Input: nums = [-2,0,1,3], and target = 2
	Output: 2
	Explanation: Because there are two triplets which sums are less than 2:
		[-2,0,1]
		[-2,0,3]

	Follow up: Could you solve it in O(n2) runtime?
*/

// tags: sort, two pointer
