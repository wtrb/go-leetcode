package product

import (
	"sort"
)

func maximumProduct(nums []int) int {
	// return useSort(nums)
	return useSingleScan(nums)
}

// Time complexity: O(n*logn). Sorting the nums array takes n*nlogn time.
// Space complexity: O(1).
func useSort(nums []int) int {
	sort.Ints(nums)
	p1 := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	p2 := nums[0] * nums[1] * nums[len(nums)-1]

	if p1 > p2 {
		return p1
	}
	return p2
}

// Time complexity: O(n)
// Space complexity: O(1)
func useSingleScan(nums []int) int {
	min1, min2 := 1<<63-1, 1<<63-1
	max1, max2, max3 := -1<<63, -1<<63, -1<<63

	for _, n := range nums {
		if n <= min1 {
			min1, min2 = n, min1
		} else if n <= min2 {
			min2 = n
		}

		if n >= max1 {
			max1, max2, max3 = n, max1, max2
		} else if n >= max2 {
			max2, max3 = n, max2
		} else if n >= max3 {
			max3 = n
		}
	}

	p1 := min1 * min2 * max1
	p2 := max3 * max2 * max1

	if p1 > p2 {
		return p1
	}
	return p2
}

// Maximum Product of Three Numbers
// https://leetcode.com/problems/maximum-product-of-three-numbers/

// tags: sort, product, single scan
