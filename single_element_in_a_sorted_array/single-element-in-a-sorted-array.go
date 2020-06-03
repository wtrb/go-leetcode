package single

// Time complexity: O(logN)
// Space complexity: O(1)
func singleNonDuplicate(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if mid&1 == 0 { // mid is even
			if nums[mid] != nums[mid+1] {
				right = mid

			} else {
				left = mid + 2
			}

		} else { // mid is odd
			if nums[mid] != nums[mid-1] {
				right = mid - 1

			} else {
				left = mid + 1
			}
		}
	}

	return nums[left]
}

// Single Element in a Sorted Array

// https://leetcode.com/problems/single-element-in-a-sorted-array/
// https://leetcode.com/submissions/detail/338326664/?from=/explore/featured/card/may-leetcoding-challenge/535/week-2-may-8th-may-14th/3327/

// https://www.youtube.com/watch?v=aFXhs190zeg

// tags: binary search
