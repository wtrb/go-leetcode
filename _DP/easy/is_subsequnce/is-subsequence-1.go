package string

// Time complexity: O(n) with n is the length of string t.
// Space complexity: O(1)
func isSubsequence1(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	count := 0
	i := 0

	for j := 0; j < len(t); j++ {
		if t[j] == s[i] {
			count++
			i++
		}

		if count == len(s) {
			return true
		}
	}

	return false
}

// Is Subsequence
// https://leetcode.com/problems/is-subsequence/

// tags: dp, dynamic programming
