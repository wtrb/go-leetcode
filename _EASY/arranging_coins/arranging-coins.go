package coin

func arrangeCoins(n int) int {
	// return loop(n)
	return binarySearch(n)
}

func loop(n int) int {
	k := 0
	for k*(k+1)/2 <= n {
		k++
	}
	return k - 1
}

// Time complexity: O(logN)
// Space complexity: O(1)
func binarySearch(n int) int {
	left, right := 0, n
	for left <= right {
		k := left + (right-left)/2
		cur := k * (k + 1) / 2

		if cur == n {
			return k
		}

		if cur < n {
			left = k + 1

		} else {
			right = k - 1
		}
	}
	return right
}

// Arranging Coins
// https://leetcode.com/problems/arranging-coins/

/*
k-th row has k coins
=> each rows has: 1coin + 2coins + 3coins + ... + kcoins
=> total coins: k*(k+1)/2

=> find k such that k*(k+1)/2 <= n
*/

// tags: binary search
