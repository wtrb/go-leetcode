package string

// Time complexity: O(m*n)
// Space complexity: O(1)
func isSubsequence0(s string, t string) bool {
	start := 0
	for i := 0; i < len(s); i++ {
		if j := index(t, s[i], start); j < 0 {
			return false

		} else {
			start = j + 1
		}
	}

	return true
}

func index(t string, c byte, start int) int {
	if start >= len(t) {
		return -1
	}

	for i := start; i < len(t); i++ {
		if t[i] == c {
			return i
		}
	}
	return -1
}

// Is Subsequence
// https://leetcode.com/problems/is-subsequence/

// tags: dp, dynamic programming
