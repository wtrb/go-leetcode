package square

// Time complexity: O(m*n)
// Space complexity: O(1)
func countSquares(matrix [][]int) int {
	count := 0

	for row := 0; row < len(matrix); row++ {
		if matrix[row][0] == 1 {
			count++
		}
	}
	for col := 1; col < len(matrix[0]); col++ { // start at 1 to avoid counting element [0][0] twice
		if matrix[0][col] == 1 {
			count++
		}
	}
	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[0]); col++ {
			if matrix[row][col] == 1 {
				squares := 1 + minInts(matrix[row][col-1], matrix[row-1][col], matrix[row-1][col-1])
				matrix[row][col] = squares
				count += squares
			}
		}
	}

	return count
}

func minInts(ints ...int) int {
	min := 1<<63 - 1
	for _, v := range ints {
		if v < min {
			min = v
		}
	}
	return min
}

// Count Square Submatrices with All Ones
// https://leetcode.com/problems/count-square-submatrices-with-all-ones/

// https://leetcode.com/submissions/detail/342874058/?from=/explore/featured/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3336/
// https://www.youtube.com/watch?v=PlZDg1VoKDg

// tags: dp, dynamic programming
