package array

// Time complexity: O(n). The entire array is traversed only once.
// Space complexity: O(n). Maximum size of the HashMap will be n, if all the elements are either 1 or 0.
func findMaxLength(nums []int) int {
	maxLen := 0
	counts := map[int]int{
		0: -1,
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count--

		} else {
			count++
		}

		if idx, ok := counts[count]; ok {
			maxLen = maxInts(maxLen, i-idx)

		} else {
			counts[count] = i
		}
	}
	return maxLen
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Contiguous Array
// https://leetcode.com/problems/contiguous-array/

/*
In this approach, we make use of a count variable, which is used to store the relative number of ones and zeros encountered so far while traversing the array. The count variable is incremented by one for every 1 encountered and the same is decremented by one for every 0 encountered.

We start traversing the array from the beginning. If at any moment, the count becomes zero, it implies that we've encountered equal number of zeros and ones from the beginning till the current index of the array(ii). Not only this, another point to be noted is that if we encounter the same count twice while traversing the array, it means that the number of zeros and ones are equal between the indices corresponding to the equal count values. The following figure illustrates the observation for the sequence [0 0 1 0 0 0 1 1]:

Index:	0	1	2	3	4	5	6	7
Array:	[0	0	1	0	0	0	1	1]
Count:	-1	-2	-1	-2	-3	-4	-3	-2
*/

// https://leetcode.com/submissions/detail/345213591/?from=/explore/featured/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3341/

// tags: hash map, map
