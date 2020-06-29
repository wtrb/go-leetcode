package pandindrome

func countSubstrings(s string) int {
	return loop(s)
}

// Time complexity: O(n^2)
// Space complexity: O(1)
func loop(s string) int {
	count := len(s)

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if ok := isPalindrome(s[i : j+1]); ok {
				count++
			}
		}
	}

	return count
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// Palindromic Substrings
// https://leetcode.com/problems/palindromic-substrings/

// tags: dp, dynamic progamming
