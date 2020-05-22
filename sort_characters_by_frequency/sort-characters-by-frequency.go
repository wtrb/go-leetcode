package charfreq

import (
	"bytes"
)

// Time complexity: O(N)
// Space complexity: O(N)
func frequencySort(s string) string {
	counter := make(map[byte]int)
	for i := range s {
		counter[s[i]]++
	}

	buckets := make([][]byte, len(s)+1)
	for k, v := range counter {
		buckets[v] = append(buckets[v], k)
	}

	result := make([]byte, 0, len(s))
	for i := len(buckets) - 1; i >= 1; i-- {
		if len(buckets[i]) > 0 {
			for _, ch := range buckets[i] {
				result = append(result, bytes.Repeat([]byte{ch}, i)...)
			}
		}
	}

	return string(result)
}

// Sort Characters By Frequency
// https://leetcode.com/problems/sort-characters-by-frequency/

// https://leetcode.com/submissions/detail/343008055/?from=/explore/featured/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3337/

// tags: bucket sort, sort, map
