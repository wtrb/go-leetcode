package string

import "strings"

// Horizontal scanning
// Time complexity: O(s)
// Space complexity: O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]

			if len(prefix) == 0 {
				return ""
			}
		}
	}

	return prefix
}

// Longest Common Prefix
// https://leetcode.com/problems/longest-common-prefix/

// tags: horizontal scanning, trie, prefix trie
