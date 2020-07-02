package number

// Time Complexity: O(log10(x))
// Space Complexity: O(1)
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	inverted := 0
	for x > inverted {
		inverted = inverted*10 + x%10
		x /= 10
	}

	if x == inverted || x == inverted/10 {
		return true
	}

	return false
}

// Palindrome Number
// https://leetcode.com/problems/palindrome-number/
