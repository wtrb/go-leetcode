package number

import "sort"

func missingNumber(nums []int) int {
	// return useMath(nums)
	return useSort(nums)
}

// Time complexity: O(n)
// Space complexity: O(1)
func useMath(nums []int) int {
	actualSum := 0
	for _, n := range nums {
		actualSum += n
	}

	expectedSum := len(nums) * (len(nums) + 1) / 2

	return expectedSum - actualSum
}

// Time complexity: O(n*logn)
// Space complexity: O(1) or (n)
func useSort(nums []int) int {
	sort.Ints(nums)

	if nums[0] != 0 {
		return 0
	}
	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			return nums[i-1] + 1
		}
	}

	return -1
}

// Missing Number
// https://leetcode.com/problems/missing-number/

// tags: xor, math, sort
