package search

func search(nums []int, target int) int {
	return binarySearch(nums, target)
}

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1

		} else {
			right = mid
		}
	}

	start := left
	left = 0
	right = len(nums) - 1

	if target >= nums[start] && target <= nums[right] {
		left = start

	} else {
		right = start
	}

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid

		} else if nums[mid] < target {
			left = mid + 1

		} else {
			right = mid - 1
		}
	}

	return -1
}

// Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/
// https://leetcode.com/submissions/detail/327454134/?from=/explore/featured/card/30-day-leetcoding-challenge/530/week-3/3304/

// tags: binary search
