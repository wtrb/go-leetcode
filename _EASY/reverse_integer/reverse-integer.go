package integer

import "math"

// Time Complexity: O(log10(x)). There are roughly log10(x) digits in x.
// Space Complexity: O(1)
func reverse(x int) int {
	result := 0
	for x != 0 {
		remainder := x % 10

		// 32 bits integer range: -2147483648 to 2147483647
		if (result > math.MaxInt32/10) || (result == math.MaxInt32/10 && remainder > 7) {
			return 0
		}
		if (result < math.MinInt32/10) || (result == math.MinInt32/10 && remainder < -8) {
			return 0
		}

		result = result*10 + remainder
		x /= 10
	}
	return result
}

// Reverse Integer
// https://leetcode.com/problems/reverse-integer/
