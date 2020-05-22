package square

// Time complexity: O(n*m)
// Space complexity: O(1)
func countSquares(matrix [][]int) int {
	count := 0

	for col := 0; col < len(matrix[0]); col++ {
		if matrix[0][col] == 1 {
			count++
		}
	}
	for row := 1; row < len(matrix); row++ {
		if matrix[row][0] == 1 {
			count++
		}
	}

	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[row]); col++ {
			if matrix[row][col] != 0 {
				squares := 1 + minInts(matrix[row][col-1], minInts(matrix[row-1][col-1], matrix[row-1][col]))
				matrix[row][col] = squares
				count += squares
			}
		}
	}

	return count
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Count Square Submatrices with All Ones
// https://leetcode.com/problems/count-square-submatrices-with-all-ones/

// https://leetcode.com/submissions/detail/342874058/?from=/explore/featured/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3336/
// https://www.youtube.com/watch?v=PlZDg1VoKDg

// tags: dp, dynamic programming
