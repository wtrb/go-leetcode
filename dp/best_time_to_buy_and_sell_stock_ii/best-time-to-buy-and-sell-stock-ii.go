package stock

// Peak Valley
// Time complexity : O(n)
// Space complexity: O(1)
func maxProfitPV(prices []int) int {
	mp := 0
	valley := prices[0]
	peak := prices[0]
	i := 0

	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]

		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		mp += peak - valley
	}

	return mp
}

// Time complexity : O(n)
// Space complexity: O(1)
func maxProfit(prices []int) int {
	mp := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			mp += prices[i] - prices[i-1]
		}
	}

	return mp
}

// Best Time to Buy and Sell Stock II
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

// tags: dp
