package sqrt

func mySqrt(x int) int {
	// return bruteForce(x)
	return binarySearch(x)
}

func bruteForce(x int) int {
	r := 0
	for r*r <= x {
		r++
	}
	return r - 1
}

func binarySearch(x int) int {
	if x == 1 {
		return 1
	}
	
	left, right := 1, x/2
	for left <= right {
		mid := left + (right-left)/2

		square := mid * mid
		if square == x {
			return mid

		} else if square >= x {
			right = mid - 1

		} else if square <= x {
			left = mid + 1
		}
	}
	return right
}

// Sqrt(x)
// https://leetcode.com/problems/sqrtx/

// tags: sqrt, binary search
