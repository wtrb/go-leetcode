package intersection

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	// return useMap(nums1, nums2)
	return useSort(nums1, nums2)
}

// Time complexity: O(m+n)
// Space complexity: O(m)
func useMap(nums1 []int, nums2 []int) []int {
	var result []int
	m := make(map[int]int)

	for _, n := range nums1 {
		m[n]++
	}

	for _, n := range nums2 {
		if m[n] > 0 {
			result = append(result, n)
			m[n]--
		}
	}

	return result
}

// Time complexity: O(logm)
// Space complexity: O(1)
func useSort(nums1 []int, nums2 []int) []int {
	var result []int

	sort.Ints(nums1)
	sort.Ints(nums2)

	p1, p2 := 0, 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			result = append(result, nums1[p1])
			p1++
			p2++

		} else if nums1[p1] < nums2[p2] {
			p1++

		} else {
			p2++
		}
	}

	return result
}

// Intersection of Two Arrays II
// https://leetcode.com/problems/intersection-of-two-arrays-ii/

// tags: map, sort, two pointer
