package triangle

func minimumTotal(triangle [][]int) int {
	// return topToBottom(triangle)
	return bottomToTop(triangle)
}

// Time complexity: O(M*N)
// Space complexity: O(1)
func topToBottom(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]

			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]

			} else {
				triangle[i][j] += minInts(triangle[i-1][j-1], triangle[i-1][j])
			}
		}
	}

	return minInts(triangle[len(triangle)-1]...)
}

// Time complexity: O(M*N)
// Space complexity: O(1)
func bottomToTop(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += minInts(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
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

// Triangle
// https://leetcode.com/problems/triangle/

// tags: dp, dynamic programming, min path sum
