package intersection

// Time complexity: O(m+n)
// Space complexity: O(m)
func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	m := make(map[int]struct{})

	for _, n := range nums1 {
		m[n] = struct{}{}
	}

	for _, n := range nums2 {
		if _, ok := m[n]; ok {
			result = append(result, n)
			delete(m, n)
		}
	}

	return result
}

// Intersection of Two Arrays
// https://leetcode.com/problems/intersection-of-two-arrays/

// tags: map
