package paint

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	for i := 1; i < len(costs); i++ {
		costs[i][0] += minInts(costs[i-1][1], costs[i-1][2])
		costs[i][1] += minInts(costs[i-1][0], costs[i-1][2])
		costs[i][2] += minInts(costs[i-1][0], costs[i-1][1])
	}

	return minInts(costs[len(costs)-1][0], costs[len(costs)-1][1], costs[len(costs)-1][2])
}

func minInts(ints ...int) int {
	min := 1<<63 - 1
	for _, v := range ints {
		if v < min {
			min = v
		}
	}
	return min
}

// Paint house
// https://leetcode.com/problems/paint-house/
/*
There are a row of n houses, each house can be painted with one of the three colors: red, blue or green. The cost of painting each house with a certain color is different. You have to paint all the houses such that no two adjacent houses have the same color.

The cost of painting each house with a certain color is represented by a n x 3 cost matrix. For example, costs[0][0] is the cost of painting house 0 with color red; costs[1][2] is the cost of painting house 1 with color green, and so on... Find the minimum cost to paint all houses.
*/

// https://www.youtube.com/watch?v=fZIsEPhSBgM

// tags: dp, dynamic programming
