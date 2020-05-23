package list

// Time Complexity: O(M + N), where M,N are the lengths of A and B respectively.
// Space Complexity: O(M + N), the maximum size of the answer.
func intervalIntersection(A [][]int, B [][]int) [][]int {
	result := make([][]int, 0)

	i := 0
	j := 0
	for i < len(A) && j < len(B) {
		lo := maxInts(A[i][0], B[j][0])
		hi := minInts(A[i][1], B[j][1])
		if lo <= hi {
			result = append(result, []int{lo, hi})
		}

		if A[i][1] < B[j][1] {
			i++

		} else {
			j++
		}
	}

	return result
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Interval List Intersections
// https://leetcode.com/problems/interval-list-intersections/

// https://leetcode.com/submissions/detail/343437659/?from=/explore/featured/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3338/

// tags: two pointers
