package twosum

func twoSum(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1
	for lo < hi {
		sum := numbers[lo] + numbers[hi]

		if sum == target {
			return []int{lo + 1, hi + 1}

		} else if sum < target {
			lo++

		} else {
			hi--
		}
	}
	return nil
}

// Two Sum II - Input array is sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

/*
	numbers = [2,7,11,15]
	2+7 <= target <= 11+15
*/

// tags: loop, two pointer
