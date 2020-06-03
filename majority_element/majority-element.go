package majelem

import (
	"sort"
)

func majorityElement(nums []int) int {
	// return hashMap(nums)
	// return sorting(nums)
	return boyerMooreVoting(nums)
}

// Time complexity : O(n)
// Space complexity : O(n)
func hashMap(nums []int) int {
	count := make(map[int]int)

	for _, n := range nums {
		count[n]++

		if c := count[n]; c > len(nums)/2 {
			return n
		}
	}

	return -1
}

// Time complexity : O(n*log(n))
// Space complexity : O(1)
func sorting(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// Time complexity : O(n)
// Space complexity : O(1)
func boyerMooreVoting(nums []int) int {
	count := 0
	candidate := 0

	for _, n := range nums {
		if count == 0 {
			candidate = n

		}

		if n == candidate {
			count++

		} else {
			count--
		}
	}

	return candidate
}

// Majority Element
// https://leetcode.com/problems/majority-element/
// https://leetcode.com/submissions/detail/335193577/?from=/explore/featured/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3321/

// tags: map, sort, boyer moore voting
