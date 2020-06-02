package jewstones

// Time complexity: O(n)
// Space complexity: O(n)
func numJewelsInStones(J string, S string) int {
	jewels := make(map[rune]struct{}, len(S))
	for _, j := range J {
		jewels[j] = struct{}{}
	}

	count := 0
	for _, s := range S {
		if _, ok := jewels[s]; ok {
			count++
		}
	}

	return count
}

// Jewels and Stones
// https://leetcode.com/problems/jewels-and-stones/
// https://leetcode.com/submissions/detail/333092243/?from=/explore/featured/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3317/

// tags: map
