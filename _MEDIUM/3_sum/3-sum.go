package threesum

import (
	"sort"
)

// Time complexity: O(n^2)
// Space complexity: O(n*logn)
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // skip to avoid duplication
			continue
		}

		left, right := i+1, len(nums)-1
		target := -nums[i]
		for left < right {
			sum := nums[left] + nums[right]
			switch {
			case sum == target:
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] { // skip to avoid duplication
					left++
				}
				for left < right && nums[right] == nums[right-1] { // skip to avoid duplication
					right--
				}

				left++
				right--

			case sum < target:
				left++

			case sum > target:
				right--
			}
		}
	}

	return result
}

// 3sum
// https://leetcode.com/problems/3sum/

// tags: sort, two pointer
