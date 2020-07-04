package string

// Time complexity: O(n)
// Space complexity: O(1)
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]

		left++
		right--
	}
}

// Reverse String
// https://leetcode.com/problems/reverse-string/

// tags: loop, two pointer
