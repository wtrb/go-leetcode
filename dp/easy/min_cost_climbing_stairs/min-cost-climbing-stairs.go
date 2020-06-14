package climbing

func minCostClimbingStairs(cost []int) int {
	// return minInts(bruteForce(cost, 0), bruteForce(cost, 1))
	// return dpprOn(cost)
	return dppr(cost)
}

// Brute force
// Time complexity: O(2^N)
// Space complexity: O(1)
func bruteForce(cost []int, step int) int {
	if step > len(cost)-1 {
		return 0
	}

	return cost[step] + minInts(bruteForce(cost, step+1), bruteForce(cost, step+2))
}

// Dynamic programming
// Time complexity: O(N)
// Space complexity: O(N)
func dpprOn(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + minInts(dp[i-1], dp[i-2])
	}
	return minInts(dp[len(cost)-1], dp[len(cost)-2])
}

// Dynamic programming
// Time complexity: O(N)
// Space complexity: O(1)
func dppr(cost []int) int {
	first, second := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		first, second = second, cost[i]+minInts(first, second)
	}
	return minInts(first, second)
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Min Cost Climbing Stairs
// https://leetcode.com/problems/min-cost-climbing-stairs/

// tags: dp, dynamic programming
