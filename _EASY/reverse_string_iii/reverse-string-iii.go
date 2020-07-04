package string

import (
	"strings"
)

// Time complexity: O(n)
// Space complexity: O(n)
func reverseWords(s string) string {
	result := []string{}
	words := strings.Split(s, " ")

	for _, word := range words {
		arr := []byte(word)

		left, right := 0, len(arr)-1
		for left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}

		result = append(result, string(arr))
	}

	return strings.Join(result, " ")
}

// Reverse Words in a String III
// https://leetcode.com/problems/reverse-words-in-a-string-iii/
