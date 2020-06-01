package timediff

import (
	"strconv"
	"strings"
)

// Time complexity: O(N)
// Space complexity: O(1)
func findMinDifference(timePoints []string) int {
	bucket := make([]bool, 24*60)

	for _, t := range timePoints {
		idx := toMinutes(t)

		if ok := bucket[idx]; ok {
			return 0

		} else {
			bucket[idx] = true
		}
	}

	min := 1<<63 - 1
	first := -1
	prev := -1
	curr := -1
	for i, ok := range bucket {
		if ok {
			if prev == -1 {
				prev = i
				first = i

			} else {
				curr = i

				min = minInts(min, curr-prev)

				prev = curr
			}
		}
	}

	return minInts(min, 24*60-curr+first)
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func toMinutes(time string) int {
	hm := strings.Split(time, ":")

	h, _ := strconv.Atoi(hm[0])
	m, _ := strconv.Atoi(hm[1])

	return h*60 + m
}

// Minimum Time Difference
// https://leetcode.com/problems/minimum-time-difference/

// tags: counting sort
