package chars

import (
	"strings"
)

func commonChars(A []string) []string {
	if len(A) == 0 {
		return nil
	}
	if len(A) == 1 {
		return strings.Split(A[0], "")
	}

	var result []string
	bucket := make([]int, 26)
	
	for _, str := range A {
		for _, c := range str {
			bucket[c-'a']++
		}
	}

	return result
}

// Find Common Characters
// https://leetcode.com/problems/find-common-characters/
