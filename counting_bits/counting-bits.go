package bits

// Time complexity: O(N)
// Space complexity: O(N)
func countBits(num int) []int {
	result := []int{0}

	l := len(result)
	for l <= num {
		for i := 0; i < l; i++ {
			result = append(result, result[i]+1)
		}
		l = len(result)
	}

	return result[:num+1]
}

func countBitsXXX(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

// Counting Bits
// https://leetcode.com/problems/counting-bits/

// https://leetcode.com/submissions/detail/346113526/?from=/explore/featured/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3343/

// tags: bits
