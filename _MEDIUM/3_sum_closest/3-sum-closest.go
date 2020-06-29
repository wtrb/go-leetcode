package threesum

import (
	"math"
	"sort"
)

// Time complexity: O(n^2)
// Space complexity: O(n*logn)
func threeSumClosest(nums []int, target int) int {
	diff := math.MaxInt64

	sort.Ints(nums)

	for i := 0; i < len(nums) && diff != 0; i++ {
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]

			if abs(sum-target) < abs(diff) {
				diff = target - sum
			}

			if sum < target {
				lo++
			} else {
				hi--
			}
		}
	}

	return target - diff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 3Sum Closest
// https://leetcode.com/problems/3sum-closest/

// tags: sort, two pointer
