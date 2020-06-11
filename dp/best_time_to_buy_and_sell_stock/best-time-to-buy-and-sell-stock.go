package stock

// Brute Force
// Time complexity: O(N^2)
// Space complexity: O(1)
func maxProfitBF(prices []int) int {
	mp := 0

	for i := 0; i < len(prices)-1; i++ {
		for j := i+1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > mp {
				mp = profit
			}
		}
	}

    return mp
}

// One Pass
// Time complexity: O(N)
// Space complexity: O(1)
func maxProfit(prices []int) int {
	min := 1<<63-1
	max := 0

	for _, p := range prices {
		if p < min {
			min = p

		} else if p - min > max {
			max = p - min
		}
	}

	return max
}

// Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

// tags: dp
