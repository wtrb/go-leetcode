package string

// Time complexity : O(N)
// Space complexity : O(1)
func firstUniqChar(s string) int {
	count := make(map[rune]int)

	for _, r := range s {
		count[r]++
	}

	for i, r := range s {
		if c := count[r]; c == 1 {
			return i
		}
	}

	return -1
}

// First Unique Character in a String
// https://leetcode.com/problems/first-unique-character-in-a-string/
// https://leetcode.com/submissions/detail/334727788/?from=/explore/featured/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3320/

// tags: map
