package twosum

func twoSum(nums []int, target int) []int {
	// return bruteForce(nums, target)

	// return useMapTwoPass(nums, target)

	return useMapOnePass(nums, target)
}

// Time complexity : O(n^2)
// Space complexity : O(1)
func bruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Time complexity : O(n)
// Space complexity : O(n)
func useMapTwoPass(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		m[n] = i
	}

	for i, n := range nums {
		if j, ok := m[target-n]; ok && j != i {
			return []int{i, j}
		}
	}

	return nil
}

// Time complexity : O(n)
// Space complexity : O(n)
func useMapOnePass(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{i, j}
		} else {
			m[n] = i
		}
	}

	return nil
}

// Two Sum
// https://leetcode.com/problems/two-sum/

// tags: loop, map
