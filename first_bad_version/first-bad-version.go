package version

// Time complexity: O(logn)
// Space complexity: O(1)
func firstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		mid := left + (right-left)/2

		if ok := isBadVersion(mid); ok {
			right = mid

		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(version int) bool {
	return true
}

// First Bad Version
// https://leetcode.com/problems/first-bad-version/
// https://leetcode.com/explore/featured/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3316/
// https://leetcode.com/submissions/detail/332662714/?from=/explore/featured/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3316/

// tags: binary search
