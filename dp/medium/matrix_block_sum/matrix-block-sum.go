package matrix

// Time complexity: O(m*n)
// Space complexity: O(m*n)
func matrixBlockSum(mat [][]int, K int) [][]int {
	rows := len(mat)
	cols := len(mat[0])

	prefixSums := make([][]int, rows+1)
	for i := 0; i < rows+1; i++ {
		prefixSums[i] = make([]int, cols+1)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			prefixSums[i+1][j+1] = mat[i][j] + prefixSums[i][j+1] + prefixSums[i+1][j] - prefixSums[i][j]
		}
	}

	blockSums := make([][]int, rows)
	for i := 0; i < rows; i++ {
		blockSums[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			r1, c1 := maxInts(0, i-K), maxInts(0, j-K)           // start
			r2, c2 := minInts(rows-1, i+K), minInts(cols-1, j+K) // end

			blockSums[i][j] = prefixSums[r2+1][c2+1] - prefixSums[r2+1][c1] - prefixSums[r1][c2+1] + prefixSums[r1][c1] // end+1 - start
		}
	}

	return blockSums
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Matrix Block Sum
// https://leetcode.com/problems/matrix-block-sum/

// tags: dp, dynamic programming, prefix sums, prefix matrix sums
