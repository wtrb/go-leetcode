package probability

func nthPersonGetsNthSeat(n int) float64 {
	// return logic(n)
	return dypr(n)
}

// Time complexity: O(1)
// Space complexity: O(1)
func logic(n int) float64 {
	if n == 1 {
		return 1.0
	}
	return 0.5
}

// Time complexity: O(n^2)
// Space complexity: O(n)
func dypr(n int) float64 {
	if n == 1 {
		return 1.0
	}
	if n == 2 {
		return 0.5
	}

	dp := make([]float64, n+1)
	dp[1] = 1.0
	dp[2] = 0.5
	// third:  1/3 + 1/3*(0.5) = 0.5
	// fourth: 1/4 + 1/4*( 1/3 + 1/3*(0.5) ) + 1/4*(0.5) = 0.5
	for i := 3; i <= n; i++ {
		p := 1.0 / float64(i)
		for j := i - 1; j >= 2; j-- {
			p += 1.0 / float64(i) * dp[j]
		}
		dp[i] = p
	}
	return dp[n]
}

// Airplane Seat Assignment Probability
// https://leetcode.com/problems/airplane-seat-assignment-probability/

// https://www.mssc.mu.edu/~paul/Puzzle/boardingpasssolution.html

// tags: dp, dynamic programmingp, lost boarding pass
