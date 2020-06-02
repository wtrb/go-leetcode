package bitwise

// Time complexity: O(1)
// Space complexity: O(1)
func rangeBitwiseAnd(m int, n int) int {
	var count uint

	for m != n {
		m >>= 1
		n >>= 1

		count++
	}

	return m << count
}

// Bitwise AND of Numbers Range
// https://leetcode.com/problems/bitwise-and-of-numbers-range
// https://leetcode.com/submissions/detail/329271159/?from=/explore/featured/card/30-day-leetcoding-challenge/531/week-4/3308/
// https://www.youtube.com/watch?v=-qrpJykY2gE

// tags: bit
