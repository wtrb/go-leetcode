package set

import "sort"

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			
		}
	}
	return nil
}

// Set Mismatch
// https://leetcode.com/problems/set-mismatch/

// tags:
