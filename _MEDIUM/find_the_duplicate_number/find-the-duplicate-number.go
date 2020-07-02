package number

import "sort"

func findDuplicate(nums []int) int {
	// return useMap(nums)
	// return useSort(nums)
	return useTwoPointer(nums)
}

// Time complexity: O(n)
// Space complexity: O(n)
func useMap(nums []int) int {
	m := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return n

		} else {
			m[n] = struct{}{}
		}
	}

	return -1
}

// Time complexity: O(n*logn)
// Space complexity: O(1) or (n)
func useSort(nums []int) int {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}

	return -1
}

// Time complexity: O(n)
// Space complexity: O(1)
func useTwoPointer(nums []int) int {
	tortoise, hare := nums[0], nums[0]
	for {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]

		if tortoise == hare {
			break
		}
	}

	tortoise = nums[0]
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}

	return tortoise
}

// Find the Duplicate Number
// https://leetcode.com/problems/find-the-duplicate-number

/*
	You must not modify the array (assume the array is read only).
	You must use only constant, O(1) extra space.
	Your runtime complexity should be less than O(n2).
	There is only one duplicate number in the array, but it could be repeated more than once.
*/

// tags: map, sort, two pointer, cycle detection
