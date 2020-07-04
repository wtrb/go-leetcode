package string

// Time complexity: O(n)
// Space complexity: O(n)
func reverseStr(s string, k int) string {
	arr := []byte(s)

	for start := 0; start < len(arr); start += 2 * k {
		left, right := start, minInts(start+k-1, len(arr)-1)

		for left < right {
			arr[left], arr[right] = arr[right], arr[left]

			left++
			right--
		}
	}

	return string(arr)
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Reverse String II
// https://leetcode.com/problems/reverse-string-ii/

// tags: loop, two pointer
