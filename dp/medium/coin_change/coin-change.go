package coin

func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	// return bruteForce(coins, 0, amount)

	// memo := make([][]int, len(coins))
	// for i := 0; i < len(coins); i++ {
	// 	memo[i] = make([]int, amount+1)
	// 	for j := 0; j < amount+1; j++ {
	// 		memo[i][j] = 1<<63 - 1
	// 	}
	// }
	// return memoize(coins, 0, amount, memo)

	memo := make([]int, amount+1)
	return topDown(coins, amount, memo)
}

func bruteForce(coins []int, i int, amount int) int {
	if i == len(coins) || amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	minCost := 1<<63 - 1

	for x := 0; x <= amount/coins[i]; x++ {
		cost := bruteForce(coins, i+1, amount-x*coins[i])
		if cost != -1 {
			minCost = minInts(minCost, x+cost)
		}
	}

	if minCost == 1<<63-1 {
		return -1
	}
	return minCost
}

func memoize(coins []int, i int, amount int, memo [][]int) int {
	if i == len(coins) || amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	if memo[i][amount] != 1<<63-1 {
		return memo[i][amount]

	} else {
		minCost := 1<<63 - 1

		for x := 0; x <= amount/coins[i]; x++ {
			cost := memoize(coins, i+1, amount-x*coins[i], memo)
			if cost != -1 {
				minCost = minInts(minCost, x+cost)
			}
		}

		if minCost == 1<<63-1 {
			return -1
		}

		memo[i][amount] = minCost
		return minCost
	}
}

// Time complexity: O(S*n), where S is the amount, n is denomination count.
// Space complexity: O(S)
func topDown(coins []int, amount int, memo []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	if memo[amount] != 0 {
		return memo[amount]

	} else {
		totalCount := 1<<63 - 1
		for _, c := range coins {
			count := topDown(coins, amount-c, memo)
			if count >= 0 && count < totalCount {
				totalCount = 1 + count
			}
		}

		if totalCount == 1<<63-1 {
			totalCount = -1
		}

		memo[amount] = totalCount

		return totalCount
	}
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Coin Change
// https://leetcode.com/problems/coin-change/

// tags: dp, dynamic programming, Knapsack
