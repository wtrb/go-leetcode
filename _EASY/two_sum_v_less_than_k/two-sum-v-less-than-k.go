package twosum

import (
	"sort"
)

// Time complexity: O(n*log(n))
// Space complexity: O(1)
func twoSum(numbers []int, k int) int {
	if numbers == nil || len(numbers) < 2 {
		return -1
	}

	sort.Ints(numbers)
	if numbers[0] + numbers[1] >= k {
		return -1
	}

	min := numbers[0] + numbers[1]
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left]+numbers[right]
		if sum < k {
			if sum == k-1 {
				return sum
			}
			min = maxInts(min, sum)
			left++

		} else {
			right--
		}
	}

	return min
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Two Sum Less Than K
// https://leetcode.com/problems/two-sum-less-than-k/

/*
	Given an array A of integers and integer K, return the maximum S such that there exists i < j with A[i] + A[j] = S and S < K. If no i, j exist satisfying this equation, return -1.

	Example 1:
	Input: A = [34,23,1,24,75,33,54,8], K = 60
	Output: 58
	Explanation: 
	We can use 34 and 24 to sum 58 which is less than 60.
	Example 2:
	Input: A = [10,20,30], K = 15
	Output: -1
	Explanation: 
	In this case it's not possible to get a pair sum less that 15.

	Note:
	1 <= A.length <= 100
	1 <= A[i] <= 1000
	1 <= K <= 2000
*/

// tags: sort, two pointer
